package mcore

import (
	//"os"
	"runtime"
)

var O_NL = "\n"

// FIXME
func GetOS() string {
	//return os.Getenv("OS")
	return runtime.GOOS
}

func IsWindows() bool {
	return String(GetOS()).IsContainIgnoreCase("WINDOWS")
}

// FIXME not work in linux
func IsLinux() bool {
	return String(GetOS()).IsContainIgnoreCase("LINUX")
}

// FIXME not work in MacOS
func IsDarwin() bool {
	return String(GetOS()).IsContainIgnoreCase("DARWIN")
}
