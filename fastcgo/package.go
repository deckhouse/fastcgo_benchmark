package fastcgo

import (
	"unsafe"
)

func EmptyAsmFunction()

func SimpleCall0(function unsafe.Pointer)

func Call0(function unsafe.Pointer)
func Call2(function unsafe.Pointer, args uintptr, result uintptr)
