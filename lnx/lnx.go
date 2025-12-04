package lnx

import "github.com/ebitengine/purego"

// Keep list of managed eagerly loaded SOs
var soList []uintptr

// Keep list of managed foreign function addresses
var ProcList [4096]uintptr

// Function to eagerly load an unmanaged SO
func UMLoad(str string) (so uintptr) {
	var err error
	so, err = purego.Dlopen(str, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {panic(err)}
	return
}

// Function to eagerly load a managed SO
func Load(str string) (so uintptr) {
	so = UMLoad(str)
	soList = append(soList, so)
	return
}

// Function to lazily load an SO
func LazyLoad(str string) (so uintptr) {
	var err error
	so, err = purego.Dlopen(str, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {panic(err)}
	return
}

// Function to find a foreign function within an eagerly loaded SO
func Proc(so uintptr, str string) (procAddr uintptr) {
	var err error
	procAddr, err = purego.Dlsym(so, str)
	if err != nil {panic(err)}
	return
}

// Function to find and manage a foreign function within an eagerly loaded SO
func MProc(so uintptr, index int, str string) {
	ProcList[index] = Proc(so, str)
	return
}

// Function to find and manage a list of foreign functions within an eagerly loaded SO
func MProcs(so uintptr, index int, strs []string) {
	for i := 0; i < len(strs); i, index = i+1, index+1 {
		MProc(so, index, strs[i])
	}
	return
}


// Function to lazily load an SO, and find and manage a foreign function within it
func MLazyLoadProc(soStr string, index int, procStr string) (so uintptr) {
	so = LazyLoad(soStr)
	MProc(so, index, procStr)
	return
}


// Function to lazily load an SO, and find and manage a list of foreign functions within it
func MLazyLoadProcs(str string, index int, strs []string) (so uintptr) {
	so = LazyLoad(str)
	MProcs(so, index, strs)
	return
}

// Simplified wrapper for purego.SyscallN
func Invoke(proc uintptr, args... uintptr) (ret uintptr) {
	ret, _, _ = purego.SyscallN(
		proc,
		args...,
	)
	return
}

// Function to release all managed eagerly loaded DLLs
func ReleaseAll() {
	for i := 1; i < len(soList); i++ {purego.Dlclose(soList[i])}
}
