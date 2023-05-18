package itoa

import (
	"io"
	"testing"
	"strings"
)

func TestUp1(t *testing.T) {
	var buffer strings.Builder
	up1(100, &buffer)
	result := buffer.String()
	if result != "\x1B[100A" {
		t.Fatalf("result=%v", result)
	}
}

func TestUp2(t *testing.T) {
	var buffer strings.Builder
	up2(100, &buffer)
	result := buffer.String()
	if result != "\x1B[100A" {
		t.Fatalf("result=%v", result)
	}
}

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
