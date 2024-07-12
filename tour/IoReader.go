package main

import "io"

type StringReader struct {
	s   string
	pos int
}

func (r StringReader) Read(buf []byte) (n int, err error) {
	if r.pos >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(buf, r.s[r.pos:])
	r.pos += n
	return n, nil
}

func NewReader(s string) *StringReader {
	return &StringReader{s, 0}
}

func main() {
	r := NewReader("Hello, World!!!")
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		if err != io.EOF {
			break
		}
		println(string(buf[:n]))
	}
}
