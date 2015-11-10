package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

var location = "/rundata/confs/mlogin.conf"

func demo() {
	lines, _ := mcore.ReadFileLines(location)

	v, err := mcore.GetConfigValue(lines, "username", ":")
	fmt.Printf("%s:%v\n", v, err)

	v2, err2 := mcore.GetConfigValue(lines, "password", ":")
	fmt.Printf("%s:%v\n", v2, err2)

	v3, err3 := mcore.GetConfigValue(lines, "none", ":")
	fmt.Printf("%s:%v\n", v3, err3)
}

func demoKeyValueArray() {
	kva, _ := mcore.NewKeyValueArrayFromFile(location, ":")
	for kva.Next() {
		fmt.Printf("%s:%v\n", kva.Key(), kva.Value())
	}
}

func main() {
	//demo()
	demoKeyValueArray()
}
