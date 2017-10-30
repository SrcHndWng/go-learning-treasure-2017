package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if (n == 0) || (n == 1) {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(10))
}
