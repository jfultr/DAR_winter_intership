package main

import (
	"fmt"
	"sort"
)

func getMinMax(a, b, c, d int) (int, int) {
	ints := []int{a, b, c, d}
	sort.Ints(ints)
	return ints[0], ints[len(ints)-1]
}
func main() {
	fmt.Println(getMinMax(3, 1, 4, 5))
}
