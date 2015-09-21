package main

import (
	"fmt"
	"github.com/mabetle/mcore/mmodel"
)

type Model struct {
	Id   string
	Name string
}

func Demo(m1 interface{}) {
	fmt.Println(mmodel.GetModelId(m1))
	m2 := mmodel.AddModelUuid(m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%+v\n", m2)
}

func main() {
	m1 := &Model{Id: "1", Name: "Demo1"}
	Demo(m1)
	m2 := Model{Id: "1", Name: "Demo1"}
	Demo(m2)
}
