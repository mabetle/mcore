package mcore

import (
	"bufio"
	"fmt"
	"github.com/mabetle/mcore/mcon"
	"os"
	"strings"
)

const (
	NEW_LINE_BYTE = byte(10)
)

// ReadLine from os.Stdio
func ReadLine() string {
	r := bufio.NewReader(os.Stdin)
	result, err := r.ReadString(NEW_LINE_BYTE)
	if err != nil {
		mcon.PrintfRed("Error:%s\n", err)
	}
	result = strings.TrimSuffix(result, "\n")
	result = strings.TrimSuffix(result, "\r")
	return result
}

// ReadArgs
func ReadArgs() []string {
	s := ReadLine()
	return ParseStringToArgs(s)
}

// ReadLineWithMsg
func ReadLineWithMsg(msgs ...interface{}) string {
	msg := fmt.Sprint(msgs...)
	if !String(msg).IsEndWith(":") {
		msg = msg + ":"
	}
	mcon.PrintGreen(msg)
	return ReadLine()
}

func ReadNotBlankLine() (result string) {
	for {
		result = ReadLine()
		if String(result).IsBlank() {
			mcon.PrintRed("input blank line, try again:")
		} else {
			break
		}
	}
	return
}

func ReadNotBlankLineWithMsg(msgs ...interface{}) string {
	msg := fmt.Sprint(msgs...)
	if !String(msg).IsEndWith(":") {
		msg = msg + ":"
	}
	fmt.Print(msg)
	return ReadNotBlankLine()
}

// ReadInt
func ReadInt(msg ...interface{}) int {
	v := ReadLineWithMsg(msg...)
	n, err := StrToInt(v)
	if err != nil {
		return ReadInt("Wrong int format,try again:")
	}
	return n
}

func ReadNotZeroInt(msg ...interface{}) int {
	v := ReadInt(msg...)
	if v == 0 {
		return ReadNotZeroInt("Input not zero int, try again:")
	}
	return v
}

// ReadBool
func ReadBool(dft bool, msg ...interface{}) bool {
	v := ReadLineWithMsg(fmt.Sprint(msg...))
	if String(v).IsBlank() {
		return dft
	}
	return String(v).ToBool()
}

// ReadSelectArray returns select array elements.
// vs should has elements
func ReadSelectArray(vs []string, msg ...interface{}) string {
	for k, v := range vs {
		fmt.Printf("%d:%s\n", k, v)
	}
	mmsg := fmt.Sprint(msg...)
	if mmsg == "" {
		mmsg = "Please select one"
	}
	s := ReadInt(mmsg)
	n := len(vs)
	// out of range
	if s > n-1 || s < 0 {
		fmt.Printf("%d out of max range, shoud from 0 to %d\n", s, n-1)
		return ReadSelectArray(vs, msg...)
	}
	return vs[s]
}
