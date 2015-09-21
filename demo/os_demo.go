package demo

import (
	"github.com/mabetle/mcore"
	"runtime"
	"fmt"
)


func init(){
	AddFunc(ShowOS, "showos")
	AddFunc(JudgeOS, "judgeos")
}

func ShowOS(){
	fmt.Printf("%v\n", runtime.GOOS)
}

func JudgeOS(){
	w:=mcore.IsWindows()
	fmt.Printf("Is Win:%v\n", w)
}
