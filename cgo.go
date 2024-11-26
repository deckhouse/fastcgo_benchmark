package cgo

/*
void empty_function() {}

int function(int a) {
  return a;
}

void fastcgo_function(void *args, void *result) {
  struct Arguments {
    int a;
  };
  struct Result {
    int value;
  };

  ((struct Result*)result)->value = ((struct Arguments*)args)->a;
}
*/
import "C"

import (
	"golang.conf/fastcgo"
	"unsafe"
)

func CGoEmptyFunctionCall() {
	C.empty_function()
}

func CGoFunctionCall() {
	C.function(5)
}

func FastCGoEmptyFunctionCall() {
	fastcgo.SimpleCall0(C.empty_function)
}

func FastCGoEmptyFunctionCallOnSystemStack() {
	fastcgo.Call0(C.empty_function)
}

func FastCGoFunctionCallOnSystemStack() {
	var args = struct {
		a int32
	}{1}

	var result = struct {
		value int32
	}{}

	fastcgo.Call2(C.fastcgo_function, uintptr(unsafe.Pointer(&args)), uintptr(unsafe.Pointer(&result)))
}

//go:noinline
func GoEmptyFunctionCall() {
}

//go:noinline
func GoFunctionCall(a int) int {
	return a
}
