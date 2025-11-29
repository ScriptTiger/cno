package cno

import (
	"syscall"
	"unsafe"
)

// Function to convert a Go string to an ANSI character array and return its pointer
func CStrA(str string) (uintptr) {
	ptr, _ := syscall.BytePtrFromString(str)
	return uintptr(unsafe.Pointer(ptr))
}

// Function to convert a Go string to a wide character array and return its pointer
func CStrW(str string) (uintptr) {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}
