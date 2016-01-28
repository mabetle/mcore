package mcore

import (
	"fmt"
)

type Results struct {
	Errors []error  //strore error
	Msgs   []string //store not error message.
}

func NewResults() *Results {
	return &Results{
		Errors: []error{},
		Msgs:   []string{},
	}
}

func (r *Results) RecordMsg(msg string) {
	r.Msgs = append(r.Msgs, msg)
}

func (r *Results) RecordError(err error) {
	r.Errors = append(r.Errors, err)
}

func (r *Results) RecordErrMsg(errMsg string) {
	r.RecordError(fmt.Errorf(errMsg))
}

func (r *Results) RecordfMsg(format string, args ...interface{}) {
	r.RecordfMsg(fmt.Sprintf(format, args...))
}

func (r *Results) RecordfError(format string, args ...interface{}) {
	r.RecordErrMsg(fmt.Sprintf(format, args...))
}

func (r *Results) IsHasError() bool {
	return len(r.Errors) > 0
}

func (r *Results) IsHasMsg() bool {
	return len(r.Msgs) > 0
}

func (r *Results) PrintErrors() {
	if !r.IsHasError() {
		fmt.Println("No errors")
		return
	}
	for i, err := range r.Errors {
		fmt.Printf("Error %d: %v\n", i+1, err)
	}
}

func (r *Results) PrintMsgs() {
	if !r.IsHasMsg() {
		fmt.Printf("No messages.\n")
	}
	for i, v := range r.Msgs {
		fmt.Printf("Msg %d: %v\n", i+1, v)
	}
}

func (r *Results) Print() {
	r.PrintErrors()
	r.PrintMsgs()
}
