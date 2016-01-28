package main

import (
	"github.com/mabetle/mcore"
)

func Demo() {
	results := mcore.NewResults()
	for i := 0; i < 10; i++ {
		results.RecordErrorMsg("error")
		results.RecordMsg("msg")
	}

	results.Print()
}

func main() {
	Demo()
}
