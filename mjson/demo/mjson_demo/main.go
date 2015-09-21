package main

import (
	"fmt"
	"mabetle/apps/web/app/mweb"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mjson"
	"reflect"
)

type DemoBaseBase struct {
	BaseName string
}

type DemoBase struct {
	DemoBaseBase
	Name string
}

type Demo struct {
	DemoBase
	Name string // overide DemoBase Name
	Age  int    `xorm:"" json:"-"`
	age  int
}

func DemoOne() {
	m := Demo{}
	m.Name = "demo"
	m.Age = 20

	str := mjson.Marshal(m)
	println(str)

	var um = Demo{}
	mjson.Unmarshal(str, &um)
	println(um.Name)
}

func DemoArray() {
	var ms []Demo = []Demo{}
	for i := 0; i < 20; i++ {
		m := Demo{}
		m.DemoBase.Name = fmt.Sprintf("demo base %v", i)
		m.Name = fmt.Sprintf("demo%v", i)
		m.Age = i
		ms = append(ms, m)
	}
	MarshalInterface(ms)
}

func DemoInterface() {
	var is []interface{} = []interface{}{}
	for i := 0; i < 20; i++ {
		m := Demo{}
		m.DemoBase.Name = fmt.Sprintf("demo base %v", i)
		m.Name = fmt.Sprintf("demo%v", i)
		m.Age = i
		is = append(is, m)
	}
	page := mweb.NewDefaultPage(is, 1)
	//fmt.Println(page.PageRowsBuildJson())
	fmt.Println(page.PageRowsJson())
	//fmt.Println(page.PageRowsFmtJson())
	//fmt.Println(page.RowsJson())
}

func MarshalInterface(rows interface{}) {
	str := mjson.Marshal(rows)
	fmt.Println(str)
}

func DemoReflect(m interface{}) {
	fs := mcore.GetFields(m)
	fmt.Println(fs)
}

func DemoEmmber(m interface{}) {
	typ := reflect.TypeOf(m)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		fmt.Printf("Not a struct value %v\n", typ)
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		// skip
		ftyp := f.Type
		//fmt.Printf("Name:%s, Type: %+v\n", f.Name, ftyp)

		if ftyp.Kind() == reflect.Struct {
			fmt.Printf("Is struct field: %s\n", f.Name)
			DemoEmmber(reflect.New(f.Type).Interface())
		} else {
			fmt.Printf("%s\n", f.Name)
		}
	}
}

func DemoCall() {
	m := Demo{}
	m.Name = "demo"
	m.DemoBase.Name = "demo base"
	m.Age = 20

	fmt.Printf("%+v", m)

}

func main() {
	//DemoArray()
	DemoInterface()
	//m := Demo{}
	//DemoReflect(m)
	//DemoEmmber(m)
	//DemoCall()
	//DemoOne()
}
