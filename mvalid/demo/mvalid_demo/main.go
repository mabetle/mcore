package main

import (
	"github.com/mabetle/mlog"
	"github.com/mabetle/mcore/mvalid"
	"github.com/mabetle/mcore/mvalid/demo"
)

var (
	logger = mlog.GetLogger("main")
)

func DemoParse() {
	m := &demo.TagModel{}
	m.Id = "1"
	m.Age = 20
	m.Name = "demo"
	mvalid.PrintValidate(m)

	wrong := demo.TagModel{}
	mvalid.PrintValidate(wrong)
}

func main() {
	//logger.SetLevel("info", "mabetle")
	logger.SetLevel("debug", "mabetle")
	DemoParse()
}
