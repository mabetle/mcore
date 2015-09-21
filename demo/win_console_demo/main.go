package main

import (
	"fmt"
	"github.com/mabetle/mcore/mcon"
)

var (
	c = mcon.NewConsole()
)

func Demo() {
	for i := 0; i <= 256; i++ {
		if i%16 == 0 {
			fmt.Println()
		}
		c.ColorPrint(fmt.Sprintf(" %d ", i), i)
	}
}

func DemoFunc() {
	c.PrintWhite("White")
	c.PrintRed("Red")
	c.PrintYellow("Yellow")
	c.PrintGreen("Green")

	c.PrintBlue("Blue")
	fmt.Println()
}

func main() {
	DemoFunc()
	Demo()
}
