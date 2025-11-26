package gui

import (
	"syscall"
	"unsafe"
)

// Windows classes

type WNDCLASSA struct {
	Style		uint32
	LpfnWndProc	uintptr
	CbClsExtra	int32
	CbWndExtra	int32
	HInstance	syscall.Handle
	HIcon		syscall.Handle
	HCursor		syscall.Handle
	HbrBackground	syscall.Handle
	LpszMenuName	*byte
	LpszClassName	*byte
}

type WNDCLASSW struct {
	Style		uint32
	LpfnWndProc	uintptr
	CbClsExtra	int32
	CbWndExtra	int32
	HInstance	syscall.Handle
	HIcon		syscall.Handle
	HCursor		syscall.Handle
	HbrBackground	syscall.Handle
	LpszMenuName	*uint16
	LpszClassName	*uint16
}

type POINT struct {
	X	int32
	Y	int32
}

type MSG struct {
	Hwnd	syscall.Handle
	Message	uint32
	WParam	uintptr
	LParam	uintptr
	Time	uint32
	Pt	POINT
}

type RECT struct {
	Left	int32
	Top	int32
	Right	int32
	Bottom	int32
}

type PAINTSTRUCT struct {
    Hdc		syscall.Handle
    FErase	int32
    RcPaint	RECT
    FRestore	int32
    FIncUpdate	int32
}

type OPENFILENAMEA struct {
	LStructSize		uint32
	HwndOwner		syscall.Handle
	HInstance		syscall.Handle
	LpstrFilter		*byte
	LpstrCustomFilter	*byte
	NMaxCustFilter		uint32
	NFilterIndex		uint32
	LpstrFile		*byte
	NMaxFile		uint32
	LpstrFileTitle		*byte
	NMaxFileTitle		uint32
	LpstrInitialDir		*byte
	LpstrTitle		*byte
	Flags			uint32
	NFileOffset		uint16
	NFileExtension		uint16
	LpstrDefExt		*byte
	LCustData		uintptr
	LpfnHook		uintptr
	LpTemplateName		*byte
	PvReserved		unsafe.Pointer
	DwReserved		uint32
	FlagsEx			uint32
}

type OPENFILENAMEW struct {
	LStructSize		uint32
	HwndOwner		syscall.Handle
	HInstance		syscall.Handle
	LpstrFilter		*uint16
	LpstrCustomFilter	*uint16
	NMaxCustFilter		uint32
	NFilterIndex		uint32
	LpstrFile		*uint16 
	NMaxFile		uint32  
	LpstrFileTitle		*uint16
	NMaxFileTitle		uint32
	LpstrInitialDir		*uint16
	LpstrTitle		*uint16
	Flags			uint32
	NFileOffset		uint16
	NFileExtension		uint16
	LpstrDefExt		*uint16
	LCustData		uintptr
	LpfnHook		uintptr 
	LpTemplateName		*uint16
	PvReserved		unsafe.Pointer 
	DwReserved		uint32
	FlagsEx			uint32
}

type BROWSEINFOA struct {
	HwndOwner		syscall.Handle
	PidlRoot		uintptr
	PszDisplayName		*byte
	LpszTitle		*byte
	UlFlags			uint32
	Lpfn			uintptr
	LParam			uintptr
	IImage			int32
}

type BROWSEINFOW struct {
	HwndOwner		syscall.Handle
	PidlRoot		uintptr
	PszDisplayName		*uint16
	LpszTitle		*uint16
	UlFlags			uint32
	Lpfn			uintptr
	LParam			uintptr
	IImage			int32
}
