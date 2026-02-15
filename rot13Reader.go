package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	for i := 0; i < n; i++ {
		char := b[i]
		if (char >= 65 && char <= 77) || (char >= 97 && char <= 109) {
			char = char + 13
		} else if (char >= 78 && char <= 90) || (char >= 110 && char <= 122) {
			char = char - 13
		}
		b[i] = char
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
