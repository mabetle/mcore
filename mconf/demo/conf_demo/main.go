package main

import (
	"github.com/mabetle/mcore/mconf/demo"
	"github.com/mabetle/mcore/mconf/xml"
	"github.com/mabetle/mcore/mconf/yaml"
)

func DemoXmlConfig() {
	c := xml.GetConfig("../data/demo.xml")
	demo.DemoConfig(c)
}

func DemoYamlConfig() {
	c := yaml.GetConfig("../data/demo.yaml")
	demo.DemoConfig(c)
}

func main() {
	DemoXmlConfig()
	DemoYamlConfig()
}
