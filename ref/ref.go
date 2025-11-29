package main

import (
	"syscall"
	"time"

	// Imports low-level utilities for working with Windows API
	"github.com/ScriptTiger/cno/win"

	// Use a dot import for the gui package to make the coding context more familiar
	. "github.com/ScriptTiger/cno/win/gui"
)

// Test ease of extensibility of cno/win by defining foreign functions in different ways

const STD_OUTPUT_HANDLE = 0xfffffff5

// Find address of foreign function and store it at win.ProcList[1000] as a managed foreign function
func init() {win.MLazyProc(Kernel32, 1000, "WriteFile")}

// Define a native Go function for the managed foreign function
func WriteFile(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[1000], args...)}

// Define a native Go function for an unmanaged foreign function
func GetStdHandle(arg uintptr) (uintptr) {return win.Invoke(win.LazyProc(Kernel32, "GetStdHandle"), arg)}

// Function to test writing to standard output using foreign functions defined above and cno.CStrA string utility function (using CStr alias) for ANSI characters
func testWrite() {
	message := "Hello, world!"
	WriteFile(
		GetStdHandle(STD_OUTPUT_HANDLE),
		CStr(message),
		uintptr(len(message)),
		0,
		0,
	)
}

// Function to test calling MessageBoxW (using MessageBox alias) and using cno.CStrW string utility function (using CStr alias) for wide characters
func testMsgBox() {
	MessageBox(
		0,
		CStr("This is a test message"),
		CStr("Test"),
		0,
	)
}

// Window dimensions
const (
	WINDOW_WIDTH = 400
	WINDOW_HEIGHT = 150
)

// Window callback function to be given to CreateWindow
func proc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (uintptr) {
	switch msg {
		case WM_CREATE:
			CreateWindowEx(
				0,
				CStr("static"),
				CStr("Click the button below!"),
				WS_CHILD | WS_VISIBLE | SS_CENTER,
				5, 5, 370, 20,
				uintptr(hwnd),
				uintptr(0),
				0, 0,
			)
			bhwnd := CreateWindowEx(
				0,
				CStr("button"),
				CStr("Click Me"),
				WS_CHILD | WS_VISIBLE,
				150, 40, 100, 20,
				uintptr(hwnd),
				uintptr(1),
				0, 0,
			)
			SetFocus(bhwnd)
			pbhwnd := CreateWindowEx(
				0,
				CStr(PROGRESS_CLASS),
				0,
				WS_CHILD | WS_VISIBLE,
				10, 80, 365, 20,
				uintptr(hwnd),
				uintptr(2),
				0, 0,
			)
			SendMessage(
				pbhwnd,
				PBM_SETRANGE,
				0,
				uintptr(100<<16),
			)
			SendMessage(
				pbhwnd,
				PBM_SETSTEP,
				1,
				0,
			)
			return 0
		case WM_DESTROY:
			PostQuitMessage(0)
			return 0
		case WM_COMMAND:
			id := int(wparam & 0xffff)
			if id == 1 {
				SetWindowText(
					GetDlgItem(uintptr(hwnd), 0),
					CStr("Please wait while the dummy progress bar finishes."),
				)
				EnableWindow(
					GetDlgItem(uintptr(hwnd), 1),
					0,
				)
				threadID := GetCurrentThreadId()
				go func() {
					for i := 0; i < 100; i++ {
						time.Sleep(100 * time.Millisecond)
						SendMessage(
							GetDlgItem(uintptr(hwnd), 2),
							PBM_STEPIT,
							0, 0,
						)
					}
					SetWindowText(
						GetDlgItem(uintptr(hwnd), 0),
						CStr("Progress bar complete! Try again if you want!"),
					)
					EnableWindow(
						GetDlgItem(uintptr(hwnd), 1),
						1,
					)
					AttachThreadInput(threadID, GetCurrentThreadId(), 1)
					SetFocus(GetDlgItem(uintptr(hwnd), 1))
				}()
			}
			return 0
	}
	return DefWindowProc(
		uintptr(hwnd),
		uintptr(msg),
		wparam,
		lparam,
	)
}

// Function to test calling CreateWindow
func testWindow() {

	// Get screen dimensions

	screenWidth := GetSystemMetrics(SM_CXSCREEN)
	screenHeight := GetSystemMetrics(SM_CYSCREEN)

	// Call CreateWindow
	// CreateWindow calls a pure Go function, not the deprecated Windows API alias of the same name
	// Internally, CreateWindowEx (an alias for either CreateWindowExA or CreateWindowExW) is used
	CreateWindow(
		proc,
		COLOR_MENU + 1,
		CStr("Sample CNO Window"),
		WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_MINIMIZEBOX | WS_VISIBLE,
		(screenWidth - uintptr(WINDOW_WIDTH))/2, (screenHeight - uintptr(WINDOW_HEIGHT))/2, uintptr(WINDOW_WIDTH), uintptr(WINDOW_HEIGHT),
	)
}

func main() {
	// Switch from the default wide character support to ANSI instead for sending text to standard output using WriteFile
	EnableAnsi(true)
	testWrite()
	// Switch back to the default wide character support
	EnableAnsi(false)
	testMsgBox()
	testWindow()
}