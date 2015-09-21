package main

import (
	"fmt"
	"github.com/mabetle/mcore/mconf"
	"github.com/mabetle/mcore/mconf/demo"
	"github.com/mabetle/mcore/mconf/ini"
	"github.com/mabetle/mlog"
)

func Demo() {
	var c mconf.Config = mconf.NewConfig(ini.NewIniConfig("../data/demo.ini", "../data/demo2.ini"))
	c.Put("put", "temp put")
	demo.DemoConfig(c)
	fmt.Println(c.GetString("put"))
}

func main() {
	mlog.SetTrace()
	Demo()
}
