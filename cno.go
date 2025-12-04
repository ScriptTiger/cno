package cno

import (
	"syscall"
	"unsafe"
)

// Function to convert a Go string to a UTF8 C string and return its pointer
func CStr(str string) (uintptr) {
	ptr, err := syscall.BytePtrFromString(str)
	if err != nil {panic(err)}
	return uintptr(unsafe.Pointer(ptr))
}

// Function to identify a UTF8 C string and return a copy of its contents as a Go string
// As a copy is returned, the C string may be freed without affecting the returned Go string
func GStr(ptr uintptr) (string) {
	len := 0
	for ;; len++ {if *(*byte)(unsafe.Pointer(ptr + uintptr(len))) == 0 {break}}
	bytes := make([]byte, len)
	copy(bytes, unsafe.Slice((*byte)(unsafe.Pointer(ptr)), len))
	return string(bytes)
}
