package mcon

import (
	"github.com/mabetle/mcore/mterm"
)

// trace debug info warn error
type ColorPrinter interface {
	PrintBlack(s string) //trace
	PrintWhite(s string) //trace

	PrintRed(s string)    //error
	PrintYellow(s string) //warn
	PrintGreen(s string)  //info

	PrintBlue(s string)    //blue
	PrintMagenta(s string) //Magenta
	PrintCyan(s string)    //debug
}

var (
	c  = NewConsole()
	xc = NewXtermConsole()
)

func GetColorPrinter() ColorPrinter {
	if mterm.IsXterm() {
		return xc
	}
	return c
}

func PrintBlack(s string) { GetColorPrinter().PrintBlack(s) }
func PrintWhite(s string) { GetColorPrinter().PrintWhite(s) }

func PrintRed(s string)    { GetColorPrinter().PrintRed(s) }
func PrintYellow(s string) { GetColorPrinter().PrintYellow(s) }
func PrintGreen(s string)  { GetColorPrinter().PrintGreen(s) }

func PrintBlue(s string)    { GetColorPrinter().PrintBlue(s) }
func PrintMagenta(s string) { GetColorPrinter().PrintMagenta(s) }
func PrintCyan(s string)    { GetColorPrinter().PrintCyan(s) }
