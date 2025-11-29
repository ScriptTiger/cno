package gui

import (
	"syscall"
	"unsafe"

	"github.com/ScriptTiger/cno"
	"github.com/ScriptTiger/cno/win"
)

// Create constants for ProcList blocks
const (
	KERNEL32_ADDR		= 0
	USER32_ADDR		= 100
	COMDLG32_ADDR		= 200
	SHELL32_ADDR		= 300
	OLE32_ADDR		= 400
	GDI32_ADDR		= 500
	ALIAS_ADDR		= 600
	KERNEL32_ALIAS_ADDR	= ALIAS_ADDR
	USER32_ALIAS_ADDR	= KERNEL32_ALIAS_ADDR+10
	COMDLG32_ALIAS_ADDR	= USER32_ALIAS_ADDR+20
	SHELL32_ALIAS_ADDR	= COMDLG32_ALIAS_ADDR+10
)

var (
	// Windows DLLs

	Kernel32		*syscall.LazyDLL
	User32			*syscall.LazyDLL
	Comctl32		*syscall.LazyDLL
	Comdlg32		*syscall.LazyDLL
	Shell32			*syscall.LazyDLL
	Ole32			*syscall.LazyDLL
	Gdi32			*syscall.LazyDLL

	// Runtime environment

	ansi			bool
)

// Windows managed foreign functions

func GetModuleHandleA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ADDR], args...)}
func GetModuleHandleW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ADDR+1], args...)}
func GetCurrentThreadId(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ADDR+2], args...)}
func WideCharToMultiByte(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ADDR+3], args...)}
func MultiByteToWideChar(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ADDR+4], args...)}

func MessageBoxA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR], args...)}
func MessageBoxW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+1], args...)}
func RegisterClassA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+2], args...)}
func RegisterClassW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+3], args...)}
func UnregisterClassA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+4], args...)}
func UnregisterClassW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+5], args...)}
func CreateWindowExA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+6], args...)}
func CreateWindowExW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+7], args...)}
func ShowWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+8], args...)}
func UpdateWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+9], args...)}
func GetMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+10], args...)}
func GetMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+11], args...)}
func TranslateMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+12], args...)}
func DispatchMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+13], args...)}
func DispatchMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+14], args...)}
func PostQuitMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+15], args...)}
func LoadCursorA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+16], args...)}
func LoadCursorW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+17], args...)}
func DefWindowProcA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+18], args...)}
func DefWindowProcW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+19], args...)}
func BeginPaint(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+20], args...)}
func EndPaint(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+21], args...)}
func FillRect(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+22], args...)}
func SendMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+23], args...)}
func SendMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+24], args...)}
func SetWindowTextA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+25], args...)}
func SetWindowTextW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+26], args...)}
func GetDlgItem(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+27], args...)}
func EnableWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+28], args...)}
func InvalidateRect(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+29], args...)}
func GetSystemMetrics(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+30], args...)}
func SetFocus(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+31], args...)}
func IsWindowEnabled(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+32], args...)}
func SetForegroundWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+33], args...)}
func AttachThreadInput(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+34], args...)}
func DestroyWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+35], args...)}
func GetWindowTextLengthA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+36], args...)}
func GetWindowTextLengthW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+37], args...)}
func GetWindowTextA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+38], args...)}
func GetWindowTextW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+39], args...)}
func GetWindowThreadProcessId(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+40], args...)}
func PostThreadMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+41], args...)}
func PostThreadMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+42], args...)}
func PostMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+43], args...)}
func PostMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ADDR+44], args...)}

func GetSaveFileNameA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ADDR], args...)}
func GetSaveFileNameW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ADDR+1], args...)}
func GetOpenFileNameA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ADDR+2], args...)}
func GetOpenFileNameW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ADDR+3], args...)}

func SHBrowseForFolderA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ADDR], args...)}
func SHBrowseForFolderW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ADDR+1], args...)}
func SHGetPathFromIDListA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ADDR+2], args...)}
func SHGetPathFromIDListW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ADDR+3], args...)}
func SHSimpleIDListFromPath(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ADDR+4], args...)}

func CoTaskMemFree(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[OLE32_ADDR], args...)}

func SetBkMode(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[GDI32_ADDR], args...)}
func SetTextColor(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[GDI32_ADDR+1], args...)}
func GetStockObject(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[GDI32_ADDR+2], args...)}
func CreateSolidBrush(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[GDI32_ADDR+3], args...)}
func DeleteObject(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[GDI32_ADDR+4], args...)}

func GetModuleHandle(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[KERNEL32_ALIAS_ADDR], args...)}

func MessageBox(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR], args...)}
func RegisterClass(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+1], args...)}
func UnregisterClass(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+2], args...)}
func CreateWindowEx(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+3], args...)}
func GetMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+4], args...)}
func DispatchMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+5], args...)}
func LoadCursor(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+6], args...)}
func DefWindowProc(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+7], args...)}
func SendMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+8], args...)}
func SetWindowText(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+9], args...)}
func GetWindowTextLength(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+10], args...)}
func GetWindowText(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+11], args...)}
func PostThreadMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+12], args...)}
func PostMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[USER32_ALIAS_ADDR+13], args...)}

func GetSaveFileName(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ALIAS_ADDR], args...)}
func GetOpenFileName(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[COMDLG32_ALIAS_ADDR+1], args...)}

func SHBrowseForFolder(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ALIAS_ADDR], args...)}
func SHGetPathFromIDList(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[SHELL32_ALIAS_ADDR+1], args...)}

// Utility aliases
func CStr(str string) (ret uintptr) {
	if ansi {ret = cno.CStrA(str)
	} else {ret = cno.CStrW(str)}
	return
}

// Load libraries and find foreign functions
func init() {
	Kernel32 = win.MLazyLoadProcs("kernel32.dll", KERNEL32_ADDR, []string{
		"GetModuleHandleA",
		"GetModuleHandleW",
		"GetCurrentThreadId",
		"WideCharToMultiByte",
		"MultiByteToWideChar",
	})

	User32 = win.MLazyLoadProcs("user32.dll", USER32_ADDR, []string{
		"MessageBoxA",
		"MessageBoxW",
		"RegisterClassA",
		"RegisterClassW",
		"UnregisterClassA",
		"UnregisterClassW",
		"CreateWindowExA",
		"CreateWindowExW",
		"ShowWindow",
		"UpdateWindow",
		"GetMessageA",
		"GetMessageW",
		"TranslateMessage",
		"DispatchMessageA",
		"DispatchMessageW",
		"PostQuitMessage",
		"LoadCursorA",
		"LoadCursorW",
		"DefWindowProcA",
		"DefWindowProcW",
		"BeginPaint",
		"EndPaint",
		"FillRect",
		"SendMessageA",
		"SendMessageW",
		"SetWindowTextA",
		"SetWindowTextW",
		"GetDlgItem",
		"EnableWindow",
		"InvalidateRect",
		"GetSystemMetrics",
		"SetFocus",
		"IsWindowEnabled",
		"SetForegroundWindow",
		"AttachThreadInput",
		"DestroyWindow",
		"GetWindowTextLengthA",
		"GetWindowTextLengthW",
		"GetWindowTextA",
		"GetWindowTextW",
		"GetWindowThreadProcessId",
		"PostThreadMessageA",
		"PostThreadMessageW",
		"PostMessageA",
		"PostMessageW",
	})

	Comctl32 = win.LazyLoad("comctl32.dll")
	win.Invoke(win.LazyProc(Comctl32, "InitCommonControls"))

	Comdlg32 = win.MLazyLoadProcs("comdlg32.dll", COMDLG32_ADDR, []string{
		"GetSaveFileNameA",
		"GetSaveFileNameW",
		"GetOpenFileNameA",
		"GetOpenFileNameW",
	})

	Shell32 = win.MLazyLoadProcs("shell32.dll", SHELL32_ADDR, []string{
		"SHBrowseForFolderA",
		"SHBrowseForFolderW",
		"SHGetPathFromIDListA",
		"SHGetPathFromIDListW",
		"SHSimpleIDListFromPath",
	})

	Ole32 = win.MLazyLoadProc("ole32.dll", OLE32_ADDR, "CoTaskMemFree")

	Gdi32 = win.MLazyLoadProcs("gdi32.dll", GDI32_ADDR, []string{
		"SetBkMode",
		"SetTextColor",
		"GetStockObject",
		"CreateSolidBrush",
		"DeleteObject",
	})
}

// Function to assign a block of aliases to foreign function addresses based on current environment
func init_alias_block(suffix string, dll *syscall.LazyDLL, block int, aliases []string) {
	for i, alias := range aliases {aliases[i] = alias+suffix}
	win.MLazyProcs(dll, block, aliases)
}

// Function to assign aliases to foreign function addresses based on current environment
func Init_aliases() {
	suffix := "W"
	if ansi {suffix = "A"}

	win.MLazyProc(Kernel32, KERNEL32_ALIAS_ADDR, "GetModuleHandle"+suffix)

	init_alias_block(suffix, User32, USER32_ALIAS_ADDR, []string{
		"MessageBox",
		"RegisterClass",
		"UnregisterClass",
		"CreateWindowEx",
		"GetMessage",
		"DispatchMessage",
		"LoadCursor",
		"DefWindowProc",
		"SendMessage",
		"SetWindowText",
		"GetWindowTextLength",
		"GetWindowText",
		"PostThreadMessage",
		"PostMessage",
	})

	init_alias_block(suffix, Comdlg32, COMDLG32_ALIAS_ADDR, []string{
		"GetSaveFileName",
		"GetOpenFileName",
	})

	init_alias_block(suffix, Shell32, SHELL32_ALIAS_ADDR, []string{
		"SHBrowseForFolder",
		"SHGetPathFromIDList",
	})

}

// Function to switch between ANSI and wide character support
func EnableAnsi(ansiBool bool) {
	ansi = ansiBool
	Init_aliases()
}

// Function to return if the current environment is ANSI
func IsAnsi() (bool) {
	return ansi
}

// Function to create a window, either ANSI or Wide character support
func CreateWindow(proc func(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (uintptr), hbr syscall.Handle, windowNamePtr uintptr, style, left, top, right, bottom uintptr) { 

	// Initialize aliases if they have not been already
	if win.ProcList[ALIAS_ADDR] == 0 {Init_aliases()}

	hInstance := GetModuleHandle(0)
	classNamePtr := CStr("cnoWindowClass")
	if ansi {
		RegisterClass(uintptr(unsafe.Pointer(&WNDCLASSA{
			LpfnWndProc:	syscall.NewCallback(proc),
			HInstance:	syscall.Handle(hInstance),
			HCursor:	syscall.Handle(LoadCursor(0, IDC_ARROW)),
			HbrBackground:	hbr,
			LpszClassName:	(*byte)(unsafe.Pointer(classNamePtr)),
		})))
	} else {
		RegisterClass(uintptr(unsafe.Pointer(&WNDCLASSW{
			LpfnWndProc:	syscall.NewCallback(proc),
			HInstance:	syscall.Handle(hInstance),
			HCursor:	syscall.Handle(LoadCursor(0, IDC_ARROW)),
			HbrBackground:	hbr,
			LpszClassName:	(*uint16)(unsafe.Pointer(classNamePtr)),
		})))
	}

	hwnd := CreateWindowEx(
		0,
		classNamePtr,
		windowNamePtr,
		style,
		left, top, right, bottom,
		0,
		0,
		hInstance,
		0,
	)

	ShowWindow(hwnd, 1)
	UpdateWindow(hwnd)

	var msg MSG
	for GetMessage(uintptr(unsafe.Pointer(&msg)), 0, 0, 0) != 0 {
		TranslateMessage(uintptr(unsafe.Pointer(&msg)))
		DispatchMessage(uintptr(unsafe.Pointer(&msg)))
	}
	UnregisterClass(classNamePtr, hInstance)
}
