package main

import "golang.org/x/tour/pic"

func Pic(dy, dx int) [][]uint8 {
	rows := make([][]uint8, dy, dy)
	for i := range rows {
		column := make([]uint8, dx, dx)
		for j := range column {
			column[j] = uint8((i + j) / 2)
		}
		rows[i] = column
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
