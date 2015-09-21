package main

import (
	"fmt"
	"github.com/mabetle/mcore/mtag"
)

type User struct {
	Name string `label:"zh='Hello' en='en:How' en_US='en_US:How a you'" xorm:"varchar"`
}

func main() {
	Demo(User{})    // demo stuct
	Demo(new(User)) // demo pointer
	Demo([]User{})  // demo array
}

func Demo(m interface{}) {
	f := "Name"
	v, _ := mtag.GetTag(m, f, "label")
	xorm, _ := mtag.GetTag(m, f, "xorm")
	fmt.Printf("%s\n", v)
	fmt.Printf("%s\n", xorm)

	fmt.Println("zh:", mtag.GetLabelZH(m, f))

	fmt.Println("en:", mtag.GetLocaleLabel(m, f, "en"))
	fmt.Println("en_US:", mtag.GetLocaleLabel(m, f, "en_US"))
	fmt.Println("zh_CN:", mtag.GetLocaleLabel(m, f, "zh_CN"))

	fmt.Println("de:", mtag.GetLocaleLabel(m, f, "de"))
	fmt.Println("End of Demo\n\n")
}
