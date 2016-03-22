package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// Demo
func Demo() {
	fmt.Println(mcore.GetFlagString("mode"))
}

func main() {
	Demo()
}
