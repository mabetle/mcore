package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

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
