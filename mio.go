package mcore

import (
	"bufio"
	"fmt"
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
		fmt.Printf("Error:%s\n", err)
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
	fmt.Print(msg)
	return ReadLine()
}

func ReadNotBlankLine() (result string) {
	for {
		result = ReadLine()
		if String(result).IsBlank() {
			fmt.Print("input blank line, try again:")
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

// ReadBool
func ReadBool(dft bool, msg ...interface{}) bool {
	v := ReadLineWithMsg(fmt.Sprint(msg...))
	if String(v).IsBlank() {
		return dft
	}
	return String(v).ToBool()
}
