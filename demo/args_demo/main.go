package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// Demo demo
func Demo() {
	fmt.Println("Hello Arg:", mcore.IsHasArg("hello"))
	fmt.Println("Dev Arg:", mcore.IsHasDevArg())
	fmt.Println("Prod Arg:", mcore.IsHasProdArg())
}

// DemoParse demo
func DemoParse() {
	args := `a=b b=c 
	c=d 
	`
	fmt.Printf("c: %s\n", mcore.GetArgString("c", "", args))
}

func main() {
	DemoParse()
}
