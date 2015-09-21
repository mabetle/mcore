package mcore

import (
	"fmt"
)

type Result struct{
	Error	error //strore error
	Msg		string //store not error message.
}

func NewSuccessResult(msg string)*Result{
	return &Result{Error:nil,Msg:msg}
}

func NewErrorResult(errMsg string)*Result{
	return &Result{Error:fmt.Errorf(errMsg),
		Msg:errMsg}
}

func NewSuccessResultf(format string, args ... interface{})*Result{
	return NewSuccessResult(fmt.Sprintf(format, args ...))
}

func NewErrorResultf(format string, args ... interface{})*Result{
	return NewErrorResult(fmt.Sprintf(format, args ... ))
}

func (r Result)IsHasError()(b bool){
	if r.Error != nil{
		b = true
	}
	return
}

// reverse to IsHasError
func (r Result)IsSuccess()(b bool){
	return !r.IsHasError()
}

// Print error msg
func (r Result)CheckError(){
	if r.IsHasError(){
		fmt.Println(r.Error)
	}
}

// print success msg
func (r Result)CheckSuccess(){
	if r.IsSuccess(){
		fmt.Println(r.Msg)
	}
}

// print result msg, if error, print error msg.
func (r Result)Check(){
	r.CheckError()
	r.CheckSuccess()
}


