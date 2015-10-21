package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

func main() {
	v := ""
	fmt.Println([]byte(v))
	v2 := mcore.ReadLineWithMsg("hello")
	fmt.Println([]byte(v))
	b := v == v2

	fmt.Print(b)
}
