package constant_test

import (
	"testing"
)

const(
	Mondat = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1<< iota
	Writable
	Executable
)

func TestConstanTry(t *testing.T){
	t.Log(Mondat,Tuesday)
}

func TestConstantTry1(t *testing.T){
	a := 7 //0111
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
}