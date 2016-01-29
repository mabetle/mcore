package main

import (
	"github.com/mabetle/mcore"
)

func Demo() {
	results := mcore.NewResults()
	for i := 0; i < 100000; i++ {
		results.RecordErrorMsg("error", i)
		results.RecordMsg("msg", i)
	}
	//results.Print()
}

func main() {
	Demo()
}
