package main

import (
	"fmt"
	"github.com/mabetle/mcore/mxml"
)

type Demo struct {
	Name string
	Age  int
}

func DemoOne() {
	m := Demo{}
	m.Name = "demo"
	m.Age = 20

	str := mxml.Marshal(m)
	println(str)

	var um = Demo{}
	mxml.Unmarshal(str, &um)
	println(um.Name)
}

func DemoArray() {
	var ms []Demo = []Demo{}
	for i := 0; i < 20; i++ {
		m := Demo{}
		m.Name = fmt.Sprintf("demo%v", i)
		m.Age = i
		ms = append(ms, m)
	}
	str := mxml.Marshal(ms)
	fmt.Println(str)
}

func main() {
	DemoArray()
}
