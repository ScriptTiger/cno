package main

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/ScriptTiger/cno"
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

// Function to test writing to standard output using foreign functions defined above and cno string utility function for ANSI characters
func testWrite() {
	message := "Hello, world!"
	WriteFile(
		GetStdHandle(STD_OUTPUT_HANDLE),
		cno.StrA(message),
		uintptr(len(message)),
		0,
		0,
	)
}

// Function to test calling MessageBoxW and using cno string utility function for wide characters
func testMsgBox() {
	MessageBoxW(
		0,
		cno.StrW("This is a test message"),
		cno.StrW("Test"),
		0,
	)
}

// Window dimensions
var (
	windowWidth = 400
	windowHeight = 150
)

// Window callback function to be given to CreateWindow
func proc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (uintptr) {
	switch msg {
		case WM_CREATE:
			CreateWindowExW(
				0,
				cno.StrW("static"),
				cno.StrW("Click the button below!"),
				WS_CHILD | WS_VISIBLE | SS_CENTER,
				5, 5, 370, 20,
				uintptr(hwnd),
				uintptr(0),
				0, 0,
			)
			bhwnd := CreateWindowExW(
				0,
				cno.StrW("button"),
				cno.StrW("Click Me"),
				WS_CHILD | WS_VISIBLE,
				150, 40, 100, 20,
				uintptr(hwnd),
				uintptr(1),
				0, 0,
			)
			SetFocus(bhwnd)
			pbhwnd := CreateWindowExW(
				0,
				cno.StrW(PROGRESS_CLASS),
				0,
				WS_CHILD | WS_VISIBLE,
				10, 80, 365, 20,
				uintptr(hwnd),
				uintptr(2),
				0, 0,
			)
			SendMessageW(
				pbhwnd,
				PBM_SETRANGE,
				0,
				uintptr(100<<16),
			)
			SendMessageW(
				pbhwnd,
				PBM_SETSTEP,
				1,
				0,
			)
			return 0
		case WM_DESTROY:
			PostQuitMessage(0)
			return 0
		case WM_PAINT:
			var paintStruct PAINTSTRUCT
			FillRect(
				BeginPaint(
					uintptr(hwnd),
					uintptr(unsafe.Pointer(&paintStruct)),
				),
				uintptr(unsafe.Pointer(&paintStruct.RcPaint)),
				uintptr(COLOR_MENU + 1),
			)
			EndPaint(
				uintptr(hwnd),
				uintptr(unsafe.Pointer(&paintStruct)),
			)
			return 0
		case WM_COMMAND:
			id := int(wparam & 0xffff)
			if id == 1 {
				SetWindowTextW(
					GetDlgItem(uintptr(hwnd), 0),
					cno.StrW("Please wait while the dummy progress bar finishes."),
				)
				EnableWindow(
					GetDlgItem(uintptr(hwnd), 1),
					0,
				)
				threadID := GetCurrentThreadId()
				go func() {
					for i := 0; i < 100; i++ {
						time.Sleep(100 * time.Millisecond)
						SendMessageW(
							GetDlgItem(uintptr(hwnd), 2),
							PBM_STEPIT,
							0, 0,
						)
					}
					SetWindowTextW(
						GetDlgItem(uintptr(hwnd), 0),
						cno.StrW("Progress bar complete! Try again if you want!"),
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
		case WM_CTLCOLORSTATIC:
			hdc := syscall.Handle(wparam)
			controlHwnd := syscall.Handle(lparam)
			if controlHwnd == syscall.Handle(GetDlgItem(uintptr(hwnd), 0)) {
				SetTextColor(
					uintptr(hdc),
					uintptr(0x00000000),
				)
				return uintptr(COLOR_MENU + 1)
			}
			return 0
	}
	return DefWindowProcW(
		uintptr(hwnd),
		uintptr(msg),
		wparam,
		lparam,
	)
}

// Function to test calling CreateWindow
func testWindow() {
	screenWidth := GetSystemMetrics(SM_CXSCREEN)
	screenHeight := GetSystemMetrics(SM_CYSCREEN)

	x := (screenWidth - uintptr(windowWidth))/2
	y := (screenHeight - uintptr(windowHeight))/2

	CreateWindow(
		proc,
		false,
		COLOR_MENU + 1,
		cno.StrW("Sample CNO Window"),
		WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_MINIMIZEBOX | WS_VISIBLE,
		x, y, uintptr(windowWidth), uintptr(windowHeight),
	)
}

func main() {
	testWrite()
	testMsgBox()
	testWindow()
}