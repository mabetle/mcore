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

func DemoLen() {
	ss := "asdf"
	zz := "中文"
	demoStrLen(ss)
	demoStrLen(zz)
}

func demoStrLen(ss string) {
	s := mcore.GetFixedWidthString(ss, 10, "-", true)
	fmt.Printf("%s Len:%d StringLen: %d StringWidth: %d\n",
		s,
		len(ss),
		mcore.StringLen(ss),
		mcore.StringWidth(ss))
}

func demoConverse() {
	a := "0"
	n, err := mcore.StrToInt(a)
	fmt.Printf("%d,%v\n", n, err)
	b := mcore.ReadInt("input int")
	fmt.Printf("%v\n", b)
}

func main() {
	//DemoTime()
	//DemoExcelTime("42009") // 2015-01-05
	//DemoSubSep()
	//DemoLen()
	demoConverse()
}
