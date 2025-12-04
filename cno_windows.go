package cno

import (
	"syscall"
	"unsafe"
)

// Function to convert a Go string to a UTF16 C string and return its pointer
func CStrW(str string) (uintptr) {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}
