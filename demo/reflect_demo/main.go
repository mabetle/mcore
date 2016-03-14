package main

import (
	"fmt"
	"github.com/mabetle/mcore"
	"mabetle/libs/account/models"
	"mabetle/libs/demo/beans"
	//"reflect"
	"time"
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

func DemoGetValue() {
	bean := beans.DemoForm{}
	bean.UserName = "demo"
	bean.Age = 33
	bean.IsEnabled = true
	bean.CreateTime = time.Now()

	v := mcore.GetReflectFieldValue(bean, "CreateTime")
	fmt.Printf("%+v\n", v)

	vk := v.Kind()
	fmt.Printf("%+v\n", vk)

	typ := v.Type()
	fmt.Printf("%+v\n", typ)
	fmt.Printf("%+v\n", typ.Name())
	fmt.Printf("%+v\n", typ.String())

	if typ.Name() == "Time" {
		fmt.Printf("%v\n", v)
	}
	vi := v.Interface()

	if vc, ok := vi.(time.Time); ok {
		a := mcore.FormatTime(vc)
		fmt.Printf("%v\n", a)
	}

	//a := mcore.GetFieldValue(&bean, "CreateTime")
	//fmt.Printf("%v\n", a)
}

func main() {
	//DemoTypeName()
	DemoGetValue()
}
