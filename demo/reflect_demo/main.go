package main

import (
	"fmt"
	"mabetle/libs/account/models"
	"github.com/mabetle/mcore"
)

func DemoUsedField() {
	m := models.Account{}
	include := "UserName,Email"
	exclude := "UserName"
	r := mcore.GetUsedFields(m, include, exclude)
	fmt.Printf("%v\n", r)
}

type Temp struct {
}

func DemoTypeName() {
	fmt.Printf("%s\n", mcore.GetElementTypeName("hello"))
	fmt.Printf("%s\n", mcore.GetElementTypeName(Temp{}))
	fmt.Printf("%s\n", mcore.GetElementTypeName(&Temp{}))

	var rows []Temp = []Temp{
		Temp{},
		Temp{},
	}

	fmt.Printf("%s\n", mcore.GetElementTypeName(rows))
}

func main() {
	DemoTypeName()
}
