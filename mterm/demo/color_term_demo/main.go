package main

import (
	"fmt"
	"github.com/mabetle/mcore/mterm"
)

var (
	cols []string = []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}
)

func Demo() {
	seed := 30
	for _, col := range cols {
		fmt.Println(mterm.ColorText(seed, col))
		//mterm.PrintGreen(col)
		seed = seed + 1
	}
}

func main() {
	Demo()
}
