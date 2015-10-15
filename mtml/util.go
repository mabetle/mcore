package mtml

import (
	"bytes"
	"html/template"
)

// MergeText
func MergeText(text string, data interface{}) (string, error) {
	t := template.New("merge")
	t.Parse(text)
	out := bytes.NewBufferString("")
	if err := t.Execute(out, data); err != nil {
		return "", err
	}
	return out.String(), nil
}

// MergeFile
func MergeFile(data interface{}, file ...string) (string, error) {
	t, err := template.ParseFiles(file...)
	if err != nil {
		return "", err
	}
	out := bytes.NewBufferString("")
	if err := t.Execute(out, data); err != nil {
		return "", err
	}
	return out.String(), nil
}
