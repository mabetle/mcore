package mcore

import (
	"fmt"
	"strings"
)

// Join
func Join(sep string, args ...interface{}) string {
	return JoinArray(sep, args)
}

// JoinArray
func JoinArray(sep string, args []interface{}) string {
	n := len(args)
	if n < 1 {
		return ""
	}
	vs := make([]string, n)
	for i, arg := range args {
		vs[i] = fmt.Sprint(arg)
	}
	return strings.Join(vs, sep)
}
