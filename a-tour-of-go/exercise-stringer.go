package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (v IPAddr) String() string {
	values := make([]string, len(v), len(v))
	for i, d := range v {
		// values[i] = strconv.Itoa(int(d))
		values[i] = fmt.Sprintf("%d", d)
	}
	return strings.Join(values, ".")
}

func main() {
	fmt.Println(IPAddr{8, 8, 8, 127})
}
