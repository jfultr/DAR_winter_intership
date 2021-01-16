package main

import "fmt"

func getFibonacci(n int) int {
	var a = 0
	var b = 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {
	fmt.Println(getFibonacci(6))
}
