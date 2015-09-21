package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/mabetle/mcore/mcon"
)

/*
go get github.com/mgutz/ansi
*/

var (
	c = mcon.NewXtermConsole()
)

func Demo() {
	// colorize a string, slowest method
	msg := ansi.Color("foo", "red+b:white")
	fmt.Print(msg)

	// create a closure to avoid escape code compilation
	phosphorize := ansi.ColorFunc("green+h:black")
	msg = phosphorize("Bring back the 80s!")
	fmt.Print(msg)

	// cache escape codes and build strings manually, faster than closure
	lime := ansi.ColorCode("green+h:black")
	reset := ansi.ColorCode("reset")

	msg = lime + "Bring back the 80s!" + reset
	fmt.Print(msg)

}

func DemoWrap() {
	c.PrintBlack("Black")
	c.PrintWhite("White")

	c.PrintRed("Red")
	c.PrintGreen("Green")
	c.PrintYellow("yellow")

	c.PrintBlue("Blue")
	c.PrintCyan("Cyan")
	c.PrintMagenta("Magenta")
}

func main() {
	DemoWrap()
}
