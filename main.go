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
	w.Write([]byte{"0123456789"[n]})
}

func up2(n int, w io.Writer) {
	w.Write([]byte{27, '['})
	_up2(n, w)
	w.Write([]byte{'A'})
}
