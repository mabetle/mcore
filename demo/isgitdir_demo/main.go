package main

import (
	"github.com/mabetle/mcore"
)

func main() {
	x := 0
	b := mcore.IsGitRepo()
	if b {
		x = 1
	}
	println(x)
}
