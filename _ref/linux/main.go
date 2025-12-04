package main

import (
	"time"

	. "github.com/ScriptTiger/cno"
	. "github.com/ScriptTiger/cno/lnx/gui"
	"github.com/ebitengine/purego"
)

var (
	// GTK objects

	hwnd, labelHwnd, buttonHwnd, pbHwnd uintptr

	// Progress bar control

	progress = make(chan float64)
	running bool
)

// Callback for button "clicked" event
func buttonClick() {
	Gtk_label_set_text(labelHwnd, CStr("Please wait while the dummy progress bar finishes."))
	Gtk_widget_set_sensitive(buttonHwnd, 0)
	go func() {
		G_idle_add(purego.NewCallback(pbCallback), 0)
		for i := 0; i < 100; i++ {
			fraction := float64(i) / 100.0
			progress <- fraction
			time.Sleep(100 * time.Millisecond)
		}
		close(progress)
		progress = make(chan float64)
		Gtk_label_set_text(labelHwnd, CStr("Progress bar complete! Try again if you want!"))
		Gtk_widget_set_sensitive(buttonHwnd, 1)
		Gtk_widget_grab_focus(buttonHwnd)
	}()
	return
}

// Callback for advancing progress bar widget
func pbCallback() bool {
	select {
		case fraction, ok := <-progress:
			if !ok {
				running = false
				return false
			}
			Gtk_progress_bar_set_fraction(pbHwnd, fraction)
			return true
		default:
			return true
	}
}

func main() {

	// Initialize GTK and set up main window
	Gtk_init(0, 0)
	hwnd = Gtk_window_new(0)
	Gtk_window_set_title(hwnd, CStr("Sample CNO Window"))
	Gtk_window_set_default_size(hwnd, 400, 150)
	Gtk_window_set_resizable(hwnd, 0)

	// Create a box with vertical orientation and add it to the window
	boxHwnd := Gtk_box_new(1, 5)
	Gtk_container_add(hwnd, boxHwnd)

	// Set up widgets and add them to the box
	labelHwnd = Gtk_label_new(CStr("Click the button below!"))
	buttonHwnd = Gtk_button_new_with_label(CStr("Click Me"))
	pbHwnd = Gtk_progress_bar_new()
	Gtk_box_pack_start(boxHwnd, labelHwnd, 0, 0, 5)
	Gtk_box_pack_start(boxHwnd, buttonHwnd, 0, 0, 5)
	Gtk_box_pack_start(boxHwnd, pbHwnd, 0, 0, 5)

	// Connect event signals to callbacks
	G_signal_connect_data(buttonHwnd, CStr("clicked"), purego.NewCallback(buttonClick), 0, 0, 0)
	G_signal_connect_data(hwnd, CStr("destroy"), purego.NewCallback(Gtk_main_quit), 0, 0, 0)

	// Show everything, set focus on the button, and start the main loop
	Gtk_widget_show_all(hwnd)
	Gtk_widget_grab_focus(buttonHwnd)
	Gtk_main()
}