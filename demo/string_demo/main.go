package main

import (
	"github.com/mabetle/mcore"
	"time"
	"fmt"
)




func DemoNow(){
	ds:=time.Now()
	mcore.PrintTime(ds)
	fmt.Printf("")
}

func DemoTime(){
	ds:="2015-09-06 00:00:00"
	t, _:=mcore.NewString(ds).ToTime()
	mcore.PrintTime(t)
}

func DemoExcelTime(days string){
	t:=mcore.String(days).ToExcelTime()
	mcore.PrintTime(t)
}

func main() {
	//DemoTime()
	DemoExcelTime("42009") // 2015-01-05
}
