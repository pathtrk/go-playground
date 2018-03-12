package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (int, error) {
	length := 0

	for {
		s, err := rot13.r.Read(p)
		length += s
		if err == io.EOF {
			break
		}
	}

	for i := 0; i < length; i++ {
		c := int(p[i])

		switch {
		case 'A' <= c && c <= 'M':
			c += 13
		case 'N' <= c && c <= 'Z':
			c -= 13
		case 'a' <= c && c <= 'm':
			c += 13
		case 'n' <= c && c <= 'z':
			c -= 13
		}

		p[i] = byte(c)
	}
	return length, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
