package mcore

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ideas from java StringBuffer.
type StringBuffer struct {
	*bytes.Buffer
}

// no args.
func NewStringBuffer(v ...interface{}) (sb StringBuffer) {
	buf := bytes.NewBufferString(fmt.Sprint(v...))
	sb.Buffer = buf
	return
}

func (sb StringBuffer) Append(args ...interface{}) {
	fmt.Fprint(sb.Buffer, args...)
}

func (sb StringBuffer) AppendByteArray(bs []byte) {
	sb.Append(string(bs))
}

// AppendLines
func (sb StringBuffer) AppendLines(lines []string) {
	for _, line := range lines {
		sb.AppendLine(line)
	}
}

// append "[a,b,c] as a,b,c\n"
func (sb StringBuffer) AppendArrayCSVLine(values []string) {
	for index, v := range values {
		sb.Append(v)
		if index < len(values)-1 {
			sb.Append(",")
		}
	}
	sb.Append("\n")
}

func (sb StringBuffer) AppendLine(args ...interface{}) {
	sb.Append(args...)
	sb.Append("\n")
}

func (sb StringBuffer) AppendEachLine(lines ...interface{}) {
	for _, line := range lines {
		sb.AppendLine(line)
	}
}

func (sb StringBuffer) AppendLinef(format string, args ...interface{}) {
	sb.Appendf(format, args...)
	sb.Append("\n")
}

func (sb StringBuffer) Appendf(format string, args ...interface{}) {
	fmt.Fprintf(sb.Buffer, format, args...)
}

func (sb StringBuffer) String() string {
	return fmt.Sprint(sb.Buffer)
}

// WriteTo
func (sb StringBuffer) WriteTo(dest io.Writer) (int64, error) {
	return sb.Buffer.WriteTo(dest)
}

// Print StringBuffer content to stdout.
func (sb StringBuffer) Print() {
	sb.WriteTo(os.Stdout)
}

// WriteToFile
func (sb StringBuffer) WriteToFile(file string) (int64, error) {
	dest, err := NewFileWriter(file)
	if err != nil {
		return 0, err
	}
	return sb.WriteTo(dest)
}

func (sb StringBuffer) Clear() {
	sb.Buffer = bytes.NewBufferString("")
}
