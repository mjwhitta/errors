package errors

import (
	"fmt"
	"runtime"
	"strings"
)

func getPkg() string {
	var caller int = 2
	var idx int
	var pkg string
	var pc uintptr

	pc, _, _, _ = runtime.Caller(caller)
	pkg = runtime.FuncForPC(pc).Name()

	if (pkg == "") || strings.HasPrefix(pkg, "main.") {
		return ""
	}

	idx = strings.LastIndex(pkg, "/")
	pkg = pkg[idx+1:]

	idx = strings.Index(pkg, ".")
	if idx < 0 {
		return pkg
	}

	return pkg[:idx] + ": "
}

// New will return a new error with a prefixed package name.
func New(str string) error {
	return fmt.Errorf("%s%s", getPkg(), str)
}

// Newf will return a new error from format string with a prefixed
// package name.
func Newf(format string, a ...any) error {
	return fmt.Errorf(getPkg()+format, a...)
}
