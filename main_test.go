package itoa

import (
	"io"
	"testing"
)

func BenchmarkUp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		up1(i, io.Discard)
	}
}

func BenchmarkUp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		up2(i, io.Discard)
	}
}
