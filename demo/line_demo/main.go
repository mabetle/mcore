package main

import (
	"github.com/mabetle/mcore"
	"fmt"
)


func demo(){
	l:="/rundata/confs/mlogin.conf"
	lines,_:=mcore.ReadFileLines(l)

	v,err:=mcore.GetConfigValue(lines,"username")
	fmt.Printf("%s:%v\n", v,err)

	v2,err2:=mcore.GetConfigValue(lines,"password")
	fmt.Printf("%s:%v\n", v2,err2)

	v3,err3:=mcore.GetConfigValue(lines,"none")
	fmt.Printf("%s:%v\n", v3,err3)
}

func main() {
	demo()
}
