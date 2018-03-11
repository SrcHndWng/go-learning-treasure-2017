package main

import (
	"flag"
	"fmt"

	"github.com/SrcHndWng/go-learning-treasure-2017/command-line-tool/utils"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "フィボナッチ数列の何番目かを指定")
	flag.Parse()
	fmt.Printf("Fib(%v) = %v\n", n, utils.Fib(n))
}
