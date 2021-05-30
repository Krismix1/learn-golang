package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		count, _ := counter[field]
		count++
		counter[field] = count

	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
