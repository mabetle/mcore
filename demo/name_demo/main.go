package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// Demo demo
func Demo() {
	name := "How_HOW-How_how"
	println(name)
	println(mcore.ToLabel(name))
	println(mcore.UpperCaseFirst("how are you"))
}

// DemoKey demo
func DemoKey() {
	keyStr := mcore.NewString("OrgName-label")
	//keyStr = keyStr.TrimStarts("label.", "msg.", "label-", "msg-")
	//keyStr = keyStr.TrimEnds(".label", ".msg", "-label", "-msg")
	keyStr = keyStr.ReplaceAll("label", "").ReplaceAll("msg", "")
	fmt.Println(keyStr)
	fmt.Println(mcore.ToLabel(keyStr.String()))
}

func main() {
	DemoKey()
}
