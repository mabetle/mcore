package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// DemoReadArray demo
func DemoReadArray() {
	vs := []string{
		"a",
		"b",
		"c",
	}
	v := mcore.ReadSelectArray(vs)
	fmt.Printf("%s\n", v)
}

func main() {
	DemoReadArray()
}
