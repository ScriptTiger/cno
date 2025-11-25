package cno

import (
	"syscall"
	"unsafe"
)

// Function to convert a Go string to an ANSI character array and return its pointer
func StrA(str string) (uintptr) {
	return uintptr(unsafe.Pointer(&append([]byte(str), 0)[0]))
}

// Function to convert a Go string to a wide character array and return its pointer
func StrW(str string) (uintptr) {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}
