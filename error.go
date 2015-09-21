package mcore

import (
	"fmt"
)

type Error struct {
	err error
}

func (t Error) Print() {
	fmt.Printf("%v\n", t)
}

//
func IsReturnHasError(args ...interface{}) (r bool) {
	defer func() {
		if err := recover(); err != nil {
			r = true
			return
		}
	}()
	nargs := len(args)
	if nargs < 1 {
		return
	}
	larg := args[nargs-1]

	if IsError(larg) && larg != nil {
		r = true
		return
	}
	return
}

func IsError(v interface{}) (r bool) {
	if _, ok := v.(error); ok {
		r = true
	}
	return
}
