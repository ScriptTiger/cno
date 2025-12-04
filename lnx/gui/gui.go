package gui

import (
	"github.com/ScriptTiger/cno/lnx"
	"github.com/ebitengine/purego"
)

var (
	// Shared object libraries

	Libgtk		uintptr
	Libglib		uintptr
	Libgobject	uintptr

	// Unmanaged foreign functions

	Gtk_main_quit func()
	Gtk_progress_bar_set_fraction func(progressBar uintptr, fraction float64)
)

// Managed foriegn functions

func Gtk_init(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR], args...)}
func Gtk_window_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+1], args...)}
func Gtk_widget_show(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+2], args...)}
func Gtk_main(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+3], args...)}
func Gtk_label_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+4], args...)}
func Gtk_container_add(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+5], args...)}
func Gtk_window_set_title(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+6], args...)}
func Gtk_window_set_default_size(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+7], args...)}
func Gtk_window_set_resizable(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+8], args...)}
func Gtk_entry_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+9], args...)}
func Gtk_box_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+10], args...)}
func Gtk_box_pack_start(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+11], args...)}
func Gtk_entry_get_text(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+12], args...)}
func Gtk_button_new_with_label(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+13], args...)}
func Gtk_progress_bar_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+14], args...)}
func Gtk_widget_destroy(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+15], args...)}
func Gtk_file_chooser_dialog_new(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+16], args...)}
func Gtk_dialog_run(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+17], args...)}
func Gtk_file_chooser_get_filename(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+18], args...)}
func Gtk_widget_set_size_request(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+19], args...)}
func Gtk_widget_show_all(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+20], args...)}
func Gtk_entry_set_text(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+21], args...)}
func Gtk_label_set_text(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+22], args...)}
func Gtk_widget_set_sensitive(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+23], args...)}
func Gtk_progress_bar_set_text(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+24], args...)}
func Gtk_widget_grab_focus(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGTK_ADDR+25], args...)}

func G_free(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGLIB_ADDR], args...)}
func G_idle_add(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGLIB_ADDR+1], args...)}

func G_signal_connect_data(args... uintptr) (uintptr) {return lnx.Invoke(lnx.ProcList[LIBGOBJECT_ADDR], args...)}

// Load libraries and find foreign functions
func init() {

	Libgtk = lnx.MLazyLoadProcs("libgtk-3.so", LIBGTK_ADDR, []string{
		"gtk_init",
		"gtk_window_new",
		"gtk_widget_show",
		"gtk_main",
		"gtk_label_new",
		"gtk_container_add",
		"gtk_window_set_title",
		"gtk_window_set_default_size",
		"gtk_window_set_resizable",
		"gtk_entry_new",
		"gtk_box_new",
		"gtk_box_pack_start",
		"gtk_entry_get_text",
		"gtk_button_new_with_label",
		"gtk_progress_bar_new",
		"gtk_widget_destroy",
		"gtk_file_chooser_dialog_new",
		"gtk_dialog_run",
		"gtk_file_chooser_get_filename",
		"gtk_widget_set_size_request",
		"gtk_widget_show_all",
		"gtk_entry_set_text",
		"gtk_label_set_text",
		"gtk_widget_set_sensitive",
		"gtk_progress_bar_set_text",
		"gtk_widget_grab_focus",
	})

	purego.RegisterLibFunc(&Gtk_main_quit, Libgtk, "gtk_main_quit")
	purego.RegisterLibFunc(&Gtk_progress_bar_set_fraction, Libgtk, "gtk_progress_bar_set_fraction")

	Libglib = lnx.MLazyLoadProcs("libglib-2.0.so", LIBGLIB_ADDR, []string{
		"g_free",
		"g_idle_add",
	})

	Libgobject = lnx.MLazyLoadProc("libgobject-2.0.so", LIBGOBJECT_ADDR, "g_signal_connect_data")
}
