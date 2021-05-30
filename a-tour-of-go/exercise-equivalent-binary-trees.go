// https://tour.golang.org/concurrency/7
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walker(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walker(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// because the tree is sorted, we will find the smallest value in the left branches
	// this way, the values emitted to the channel will be sorted
	walker(t.Left, ch)
	ch <- t.Value
	walker(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	values1 := []int{}
	values2 := []int{}
	// if we assume that both trees have the same number of nodes,
	// then we don't need to store values in arrays, and only compare
	// each value from the channels
	for v := range ch1 {
		values1 = append(values1, v)
	}
	for v := range ch2 {
		values2 = append(values2, v)
	}

	// because each channel will emit all values in sorted order
	// we can simplify the equality check by comparing elements at the same index
	return equal(values1, values2)
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	exists := make(map[int]bool)
	for _, v := range arr1 {
		exists[v] = true
	}
	for _, v := range arr2 {
		if !exists[v] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(4)))
}
