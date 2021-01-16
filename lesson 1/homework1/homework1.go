package main

import (
	"fmt"
	"sort"
)

func min(a, b, c, d int) int {
	ints := []int{a, b, c, d}
	sort.Ints(ints)
	return ints[0]
}
func main() {
	fmt.Println(min(3, 1, 4, 5))
}
