package main

import (
	"github.com/mabetle/mcore"
	"fmt"
)

func Demo(){
	fmt.Println(mcore.GetFlagString("mode"))
}

func main() {
	Demo()
}
