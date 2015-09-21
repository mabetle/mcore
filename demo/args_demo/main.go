package main

import (
	"github.com/mabetle/mcore"
	"fmt"
)

func main() {
	fmt.Println("Hello Arg:", mcore.IsHasArg("hello"))
	fmt.Println("Dev Arg:", mcore.IsHasDevArg())
	fmt.Println("Prod Arg:", mcore.IsHasProdArg())
}

