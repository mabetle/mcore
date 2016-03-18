package mcore

import (
	"strings"
)

type Args []string

// ParseStringToArgs line to args
// FIXME not work as expect
func ParseStringToArgs(s string) []string {
	r := []string{}
	s = strings.TrimSpace(s)
	if s == "" {
		return r
	}
	args := strings.Split(s, " ")
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		arg = strings.Trim(arg, `"`) // delete wrap "
		if arg != "" {
			r = append(r, arg)
		}
	}
	return r
}

func NewArgsFromString(s string) Args {
	return Args(ParseStringToArgs(s))
}

func NewArgs(args []string) Args {
	return Args(args)
}

func (a Args) IsHasFlag(flag string) bool {
	return String(flag).IsInArray(a)
}

func (a Args) ParseString(flag string) (r string) {

	return
}

func (a Args) ParseInt(flag string) (r int) {

	return
}

func (a Args) NArgs() int {
	return len(a)
}

func (a Args) NFlags() (r int) {

	return
}

func (a Args) VArgs(index int) string {
	return a[index]
}

// GetArgString arg format: a=b
func GetArgString(name string, defaultValue string, args ...string) string {
	// process Args
	for _, a := range args {
		arg := GetString(a)
		//arg format: a=b
		arg = strings.TrimSpace(arg)
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		if NewString(key).IsEqualIgnoreCase(name) {
			return value
		}
	}
	return defaultValue
}

// GetArgBool
func GetArgBool(name string, defaultValue bool, args ...string) bool {
	dVS := "0"
	if defaultValue {
		dVS = "1"
	}
	bStr := GetArgString(name, dVS, args...)
	return NewString(bStr).ToBool()
}

//
func GetArgInt(name string, defaultValue int, args ...string) int {
	iStr := GetArgString(name, "", args...)
	if iStr == "" {
		return defaultValue
	}
	n, err := NewString(iStr).ToInt()
	if err != nil {
		return defaultValue
	}
	return n
}

func GetArgExists(name string, args ...string) bool {
	for _, a := range args {
		arg := GetString(a)
		arg = strings.TrimSpace(arg)
		//arg format: a=b
		key := strings.Split(arg, "=")[0]
		key = strings.TrimSpace(key)
		if NewString(key).IsEqualIgnoreCase(name) {
			return true
		}
	}
	return false
}
