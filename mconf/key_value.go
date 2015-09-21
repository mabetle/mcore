package mconf

import (
	"github.com/mabetle/mcore"
)

// KeyValueConfig implements Config interface.
type KeyValueConfig struct {
	mcore.StringKeyValue // store key and value.
	KeyValueLoader       //load keys and values.
}

// KeyValueLoader
type KeyValueLoader interface {
	LoadKeyValue() mcore.StringKeyValue
}

// NewConfig
func NewConfig(loader KeyValueLoader) Config {
	c := &KeyValueConfig{}
	c.KeyValueLoader = loader
	c.StringKeyValue = c.LoadKeyValue()
	return c
}
