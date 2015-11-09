package mcore

import (
	"fmt"
)

// GetFileConfigValue returns config value from config files
func GetFileConfigValue(location string, key string)(string,error){
	lines, err:= ReadFileLines(location)
	if err!=nil{
		return "", err
	}
	return GetConfigValue(lines, key)
}

// GetConfigValue returns config value from lines.
func GetConfigValue(lines []string, key string) (string, error){
	for _, line := range lines {
		lineS := NewString(line)
		if lineS.IsStartsIgnoreCase("#","//") || lineS.IsBlank() {
			continue
		}
		var kv []string
		if lineS.IsContains(":") {
			kv = lineS.Split(":")
		}
		if lineS.IsContains("="){
			kv = lineS.Split("=")
		}
		// bad key value
		if len(kv) < 2 {
			continue
		}
		lineKey:=kv[0]
		lineValue:=kv[1]
		// found
		if NewString(lineKey).IsEqualIgnoreCase(key){
			return lineValue, nil
		}
	}
	// not found
	return "", fmt.Errorf("not found value for key: %s", key)
}

