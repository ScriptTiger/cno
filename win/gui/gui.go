package gui

import (
	"syscall"
	"unsafe"

	"github.com/ScriptTiger/cno"
	"github.com/ScriptTiger/cno/win"
)

// Windows DLLs
var (
	Kernel32		*syscall.LazyDLL
	User32			*syscall.LazyDLL
	Comctl32		*syscall.LazyDLL
	Comdlg32		*syscall.LazyDLL
	Gdi32			*syscall.LazyDLL
)

// Windows managed foreign functions

func GetModuleHandleA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[0], args...)}
func GetModuleHandleW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[1], args...)}
func GetCurrentThreadId(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[2], args...)}

func MessageBoxA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[100], args...)}
func MessageBoxW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[101], args...)}
func RegisterClassA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[102], args...)}
func RegisterClassW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[103], args...)}
func UnregisterClassA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[104], args...)}
func UnregisterClassW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[105], args...)}
func CreateWindowExA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[106], args...)}
func CreateWindowExW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[107], args...)}
func ShowWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[108], args...)}
func UpdateWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[109], args...)}
func GetMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[110], args...)}
func GetMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[111], args...)}
func TranslateMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[112], args...)}
func DispatchMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[113], args...)}
func DispatchMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[114], args...)}
func PostQuitMessage(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[115], args...)}
func LoadCursorA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[116], args...)}
func LoadCursorW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[117], args...)}
func DefWindowProcA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[118], args...)}
func DefWindowProcW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[119], args...)}
func BeginPaint(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[120], args...)}
func EndPaint(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[121], args...)}
func FillRect(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[122], args...)}
func SendMessageA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[123], args...)}
func SendMessageW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[124], args...)}
func SetWindowTextA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[125], args...)}
func SetWindowTextW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[126], args...)}
func GetDlgItem(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[127], args...)}
func EnableWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[128], args...)}
func InvalidateRect(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[129], args...)}
func GetSystemMetrics(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[130], args...)}
func SetFocus(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[131], args...)}
func IsWindowEnabled(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[132], args...)}
func SetForegroundWindow(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[133], args...)}
func AttachThreadInput(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[134], args...)}

func GetSaveFileNameA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[200], args...)}
func GetSaveFileNameW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[201], args...)}
func GetOpenFileNameA(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[202], args...)}
func GetOpenFileNameW(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[203], args...)}

func SetBkMode(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[300], args...)}
func SetTextColor(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[301], args...)}
func GetStockObject(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[302], args...)}
func CreateSolidBrush(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[303], args...)}
func DeleteObject(args... uintptr) (uintptr) {return win.Invoke(win.ProcList[304], args...)}

// Load libraries and find foreign functions
func init() {
	Kernel32 = win.LazyLoad("kernel32.dll")
	win.MLazyProcs(Kernel32, 0, []string{
		"GetModuleHandleA",
		"GetModuleHandleW",
		"GetCurrentThreadId",
	})

	User32 = win.LazyLoad("user32.dll")
	win.MLazyProcs(User32, 100, []string{
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
	})

	Comctl32 = win.LazyLoad("comctl32.dll")
	win.Invoke(win.LazyProc(Comctl32, "InitCommonControls"))

	Comdlg32 = win.LazyLoad("comdlg32.dll")
	win.MLazyProcs(Comdlg32, 200, []string{
		"GetSaveFileNameA",
		"GetSaveFileNameW",
		"GetOpenFileNameA",
		"GetOpenFileNameW",
	})

	Gdi32 = win.LazyLoad("gdi32.dll")
	win.MLazyProcs(Gdi32, 300, []string{
		"SetBkMode",
		"SetTextColor",
		"GetStockObject",
		"CreateSolidBrush",
		"DeleteObject",
	})
}

// Function to create a window, either ANSI or Wide character support
func CreateWindow(proc func(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (uintptr), ansi bool, hbr syscall.Handle, windowNamePtr uintptr, style, left, top, right, bottom uintptr) { 
	var hInstance, classNamePtr, hwnd uintptr
	if ansi {
		hInstance = GetModuleHandleA(0)
		classNamePtr = cno.StrA("cnoWindowClassA")

		RegisterClassA(uintptr(unsafe.Pointer(&WNDCLASSA{
			LpfnWndProc:	syscall.NewCallback(proc),
			HInstance:	syscall.Handle(hInstance),
			HCursor:	syscall.Handle(LoadCursorA(0, 32512)),
			HbrBackground:	hbr,
			LpszClassName:	(*byte)(unsafe.Pointer(classNamePtr)),
		})))

		hwnd = CreateWindowExA(
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
	} else {
		hInstance = GetModuleHandleW(0)
		classNamePtr = cno.StrW("cnoWindowClassW")

		RegisterClassW(uintptr(unsafe.Pointer(&WNDCLASSW{
			LpfnWndProc:	syscall.NewCallback(proc),
			HInstance:	syscall.Handle(hInstance),
			HCursor:	syscall.Handle(LoadCursorW(0, 32512)),
			HbrBackground:	syscall.Handle(COLOR_WINDOW + 1),
			LpszClassName:	(*uint16)(unsafe.Pointer(classNamePtr)),
		})))

		hwnd = CreateWindowExW(
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
	}

	ShowWindow(hwnd, 1)
	UpdateWindow(hwnd)

	var msg MSG
	if ansi {
		for GetMessageA(uintptr(unsafe.Pointer(&msg)), 0, 0, 0) != 0 {
			TranslateMessage(uintptr(unsafe.Pointer(&msg)))
			DispatchMessageA(uintptr(unsafe.Pointer(&msg)))
		}
		UnregisterClassA(classNamePtr, hInstance)
	} else {
		for GetMessageW(uintptr(unsafe.Pointer(&msg)), 0, 0, 0) != 0 {
			TranslateMessage(uintptr(unsafe.Pointer(&msg)))
			DispatchMessageW(uintptr(unsafe.Pointer(&msg)))
		}
		UnregisterClassW(classNamePtr, hInstance)
	}
}
