package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		n := b - 'A'
		return 'A' + (n+13)%26
	}
	if 'a' <= b && b <= 'z' {
		n := b - 'a'
		return 'a' + (n+13)%26
	}
	return b
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	nr, ne := rot.r.Read(b)
	if ne != nil {
		return nr, ne
	}

	for i := 0; i < nr; i++ {
		b[i] = rot13(b[i])
	}

	return nr, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
