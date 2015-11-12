package mcore

import (
	"strings"
)

type Args []string

// Parse
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
