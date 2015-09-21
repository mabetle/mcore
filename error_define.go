package mcore

import (
	"fmt"
)

type ErrorType int

const(
	TYPE_EXCEPTION ErrorType = iota
	TYPE_RUNTIME_EXCEPTION
)

// app can not go on. app should check it.
type Exception struct{
	ErrorType
	err error
}

func NewExceptionf(format string, args ... interface{})Exception{
	return NewException(fmt.Errorf(format, args ...))
}

func NewException(err error)Exception{
	return Exception{ErrorType:TYPE_EXCEPTION, err:err}
}

func (e Exception)Error()string{
	return e.err.Error()
}

func (e Exception)ErrorString()(string){
	return e.Error()
}

// depends on apps, not check.
type RuntimeException struct{
	Exception
}

func NewRuntimeExceptionf(format string, args ... interface{})RuntimeException{
	return NewRuntimeException(fmt.Errorf(format, args ...))
}

func NewRuntimeException(err error)RuntimeException{
	e:=Exception{ErrorType:TYPE_RUNTIME_EXCEPTION, err:err}
	return RuntimeException{Exception:e}
}

// RuntimeException
type FileNotFoundError struct{
	RuntimeException
}

func NewFileNotFoundError(name string)FileNotFoundError{
	e:=NewRuntimeExceptionf("File not found. File: %s ",name)
	return FileNotFoundError{RuntimeException:e}
}




func IsRuntimeException(e Exception)bool{
	return e.ErrorType == TYPE_RUNTIME_EXCEPTION
}

func IsException(e Exception)bool{
	return e.ErrorType == TYPE_EXCEPTION
}

func IsRuntimeExceptionType(err error)bool{
	if _,ok:=err.(RuntimeException);ok{
		return true
	}
	return false
}

func IsExceptionType(err error)bool{
	return !IsRuntimeExceptionType(err)
}
