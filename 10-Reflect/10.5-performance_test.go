package Reflect

import (
	"testing"
)

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(100)
	}
}

func BenchmarkRSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset(100)
	}
}

func BenchmarkRSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset2(100)
	}
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkRCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rcall()
	}
}
