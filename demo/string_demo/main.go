package main

import (
	"fmt"
	"github.com/mabetle/mcore"
	"time"
)

func DemoNow() {
	ds := time.Now()
	mcore.PrintTime(ds)
	fmt.Printf("")
}

func DemoTime() {
	ds := "2015-09-06 00:00:00"
	t, _ := mcore.NewString(ds).ToTime()
	mcore.PrintTime(t)
}

func DemoExcelTime(days string) {
	t := mcore.String(days).ToExcelTime()
	mcore.PrintTime(t)
}

func DemoSubSep() {
	in := mcore.NewString("abc?def")
	fmt.Println(in.SubLeftSep("?"))
	fmt.Println(in.SubRightSep("?"))
}

func main() {
	//DemoTime()
	//DemoExcelTime("42009") // 2015-01-05
	DemoSubSep()
}
