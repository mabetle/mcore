package mcore

import (
	"fmt"
)

func IsValueEqual(a, b interface{})bool{
	var sa, sb string
	sa = fmt.Sprint("%v", a)
	sb = fmt.Sprint("%v", b)
	if sa == sb {
		return true
	}
	return false
}

func AppendBefore(old []interface{}, value interface{})(r []interface{}){
	r = append(r, value)
	for _, v := range old {
		r = append(r, v)
	}
	return r
}

func Append(old []interface{}, value interface{})([]interface{}){
	return append(old, value)
}

