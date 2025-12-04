package win

import "syscall"

// Keep list of managed eagerly loaded DLLs
var dllList []*syscall.DLL

// Keep list of managed foreign function addresses
var ProcList [4096]uintptr

// Function to eagerly load an unmanaged DLL
func UMLoad(str string) (dll *syscall.DLL) {
	dll = syscall.MustLoadDLL(str)
	return
}

// Function to eagerly load a managed DLL
func Load(str string) (dll *syscall.DLL) {
	dll = UMLoad(str)
	dllList = append(dllList, dll)
	return
}

// Function to lazily load a DLL
func LazyLoad(str string) (dll *syscall.LazyDLL) {
	dll = syscall.NewLazyDLL(str)
	return
}

// Function to find a foreign function within an eagerly loaded DLL
func Proc(dll *syscall.DLL, str string) (procAddr uintptr) {
	procAddr = dll.MustFindProc(str).Addr()
	return
}

// Function to find and manage a foreign function within an eagerly loaded DLL
func MProc(dll *syscall.DLL, index int, str string) {
	ProcList[index] = Proc(dll, str)
	return
}

// Function to find and manage a list of foreign functions within an eagerly loaded DLL
func MProcs(dll *syscall.DLL, index int, strs []string) {
	for i := 0; i < len(strs); i, index = i+1, index+1 {
		MProc(dll, index, strs[i])
	}
	return
}

// Function to find a foreign function within a lazily loaded DLL
func LazyProc(dll *syscall.LazyDLL, str string) (procAddr uintptr) {
	procAddr = dll.NewProc(str).Addr()
	return
}

// Function to find and manage a foreign function within a lazily loaded DLL
func MLazyProc(dll *syscall.LazyDLL, index int, str string) {
	ProcList[index] = LazyProc(dll, str)
	return
}

// Function to lazily load a DLL, and find and manage a foreign function within it
func MLazyLoadProc(dllStr string, index int, procStr string) (dll *syscall.LazyDLL) {
	dll = LazyLoad(dllStr)
	ProcList[index] = LazyProc(dll, procStr)
	return
}

// Function to find and manage a list of foreign functions within a lazily loaded DLL
func MLazyProcs(dll *syscall.LazyDLL, index int, strs []string) {
	for i := 0; i < len(strs); i, index = i+1, index+1 {
		MLazyProc(dll, index, strs[i])
	}
	return
}

// Function to lazily load a DLL, and find and manage a list of foreign functions within it
func MLazyLoadProcs(str string, index int, strs []string) (dll *syscall.LazyDLL) {
	dll = LazyLoad(str)
	for i := 0; i < len(strs); i, index = i+1, index+1 {
		MLazyProc(dll, index, strs[i])
	}
	return
}

// Simplified wrapper for syscall.SyscallN
func Invoke(proc uintptr, args... uintptr) (ret uintptr) {
	ret, _, _ = syscall.SyscallN(
		proc,
		args...,
	)
	return
}

// Function to release all managed eagerly loaded DLLs
func ReleaseAll() {
	for i := 1; i < len(dllList); i++ {dllList[i].Release()}
}
