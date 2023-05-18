# エスケープシーケンスの数値を出力するときに、fmt を使わないとどれだけ効果があるか？の検証

## main.go

- up1 は fmt を使う
- up2 は自前でやる

```main.go
package itoa

import (
    "fmt"
    "io"
)

func up1(n int, w io.Writer) {
    if n == 1 {
        fmt.Fprintf(w, "\x1B[A")
    } else if n > 1 {
        fmt.Fprintf(w, "\x1B[%dA", n)
    }
}

func _up2(n int, w io.Writer) {
    if n >= 10 {
        _up2(n/10, w)
        n %= 10
    }
    w.Write([]byte{'0' + byte(n)})
}

func up2(n int, w io.Writer) {
    w.Write([]byte{27, '['})
    _up2(n, w)
    w.Write([]byte{'A'})
}
```

## main_test.go

```main_test.go
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
```

## `go test -bench . -benchmem |`

```go test -bench . -benchmem |
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-go-itoa
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkUp1-4   	 8767718	       129.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkUp2-4   	 8311158	       162.1 ns/op	       9 B/op	       8 allocs/op
PASS
ok  	github.com/hymkor/study-go-itoa	3.174s
```

結論: fmt の方が速い上に、ゼロアローケーション（これは驚き）
