package mapp

import (
	"fmt"
	"github.com/mabetle/mcore"
)

var Models = make(map[string]interface{})

func RegModels(models ...interface{}) {
	for _, model := range models {
		pkg := mcore.GetPkgPath(model)
		typ := mcore.GetTypeName(model)
		key := fmt.Sprintf("%s-%s", pkg, typ)
		Models[key] = model
	}
}
