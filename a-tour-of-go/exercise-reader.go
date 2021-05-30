package main

//  Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(buffer []byte) (int, error) {
	// https://stackoverflow.com/questions/27839140/solution-for-http-tour-golang-org-methods-11
	// for i := 0; i < len(buffer); i++ {
	// 	buffer[i] = 65
	// }
	// return len(buffer), nil
	buffer[0] = 'A'
	return 1, nil // The key take away is to return nil error
}

func ReadFromString() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	reader.Validate(MyReader{})
}
