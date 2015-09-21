package main

import (
	"github.com/mabetle/mcore"
)

func main() {
	for i := 0; i < 10; i++ {
		println(mcore.NewRandomPassword())
	}
}
