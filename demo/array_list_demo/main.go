package main

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// User demo
type User struct {
	Name string
}

func (m User) String() string {
	return m.Name
}

// Demo demo
func Demo() {
	l := mcore.NewArrayList(User{})
	l.Put(User{Name: "a"})
	l.Put(User{Name: "b"})
	fmt.Println(l.Len())
	fmt.Println(l.Type())
	l.Print()
}

// Demo2 demo
func Demo2() {
	s := ""
	l := mcore.NewArrayList(s)
	l.Put("Hello", "ABC")
	l.Print()
}

func main() {
	//Demo()
	Demo2()
}
