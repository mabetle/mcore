package mcore

import (
	"log"
)

type StringKeyValue map[string]string

func NewStringKeyValue() StringKeyValue {
	return make(map[string]string)
}

// put key value
func (c StringKeyValue) Put(key, value string) {
	c[key] = value
}

func (c StringKeyValue) IsContain(key string) bool {
	_, ok := c[key]
	return ok
}

func (c StringKeyValue) GetString(key string) string {
	return c.GetStringWithDefault(key, "")
}

func (c StringKeyValue) GetStringWithDefault(key, defaultValue string) string {
	if c.IsContain(key) {
		return c[key]
	}
	log.Printf("Error: not contains key: %s", key)
	return defaultValue
}

func (c StringKeyValue) GetInt(key string) int {
	return String(c.GetString(key)).ToIntNoError()
}

func (c StringKeyValue) GetBool(key string) bool {
	return String(c.GetString(key)).ToBool()
}
