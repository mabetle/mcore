package main

import (
	"fmt"
	"github.com/mabetle/mcore/mterm"
)

func main() {
	fmt.Println(mterm.Black("Black()"))
	fmt.Println(mterm.White("White()"))

	fmt.Println(mterm.Red("Red()"))
	fmt.Println(mterm.Green("Green()"))
	fmt.Println(mterm.Yellow("Yellow()"))

	fmt.Println(mterm.Blue("Blue()"))
	fmt.Println(mterm.Magenta("Magenta()"))
	fmt.Println(mterm.Cyan("Cyan()"))
}
