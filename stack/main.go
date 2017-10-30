package main

import (
	"fmt"

	"./stack"
)

func main() {
	ary := []string{"a", "b", "c", "d", "e", "f", "g"}
	const limit = 5
	s := stack.NewStack(limit)

	for _, v := range ary {
		s.Push(v)
	}
	fmt.Println(s)

	for i := 0; i < limit; i++ {
		poped := s.Pop()
		fmt.Println("poped = " + poped)
		fmt.Println(s)
	}
}
