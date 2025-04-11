package cgo

import (
	"testing"
)

func BenchmarkCGoEmptyFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CGoEmptyFunctionCall()
	}
}

func BenchmarkCGoFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CGoFunctionCall()
	}
}

func BenchmarkCGoNoCallbackFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CGoNoCallbackFunctionCall()
	}
}

func BenchmarkGoEmptyFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoEmptyFunctionCall()
	}
}

func BenchmarkGoFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoFunctionCall(5)
	}
}

func BenchmarkFastCGoEmptyFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastCGoEmptyFunctionCall()
	}
}

func BenchmarkFastCGoEmptyFunctionCallOnSystemStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastCGoEmptyFunctionCallOnSystemStack()
	}
}

func BenchmarkFastCGoFunctionCallOnSystemStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastCGoFunctionCallOnSystemStack()
	}
}
