package main

import (
	"io"
	"os"
	"strings"
)

// https://tour.golang.org/methods/23

type Rot13Reader struct {
	r io.Reader
}

func (r Rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err == io.EOF {
		return 0, io.EOF
	}
	for i := 0; i < n; i++ {
		letter := b[i]
		if (letter >= 'A' && letter < 'N') || (letter >= 'a' && letter < 'n') {
			letter += 13
		} else if (letter > 'M' && letter <= 'Z') || (letter > 'm' && letter <= 'z') {
			// if the letter's index is between M - Z, subtract 13 from its index
			letter -= 13
		}
		b[i] = letter
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
