package main

import (
	"flag"
	"fmt"

	"./utils"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "フィボナッチ数列の何番目かを指定")
	flag.Parse()
	fmt.Printf("Fib(%v) = %v\n", n, utils.Fib(n))
}
