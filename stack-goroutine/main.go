package main

import (
	"log"
	"sync"
	"time"

	"github.com/SrcHndWng/go-learning-treasure-2017/stack-goroutine/stack"
)

func main() {
	var wg sync.WaitGroup
	ary := []string{"a", "b", "c", "d", "e", "f", "g"}
	const limit = 5
	s := stack.NewStack(limit)
	for _, v := range ary {
		wg.Add(1)
		go func(v string) {
			time.Sleep(1 * time.Second)
			s.Push(v)
			wg.Done()
		}(v)
		wg.Wait()
	}
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func() {
			poped := s.Pop()
			log.Printf("poped = %s, stack = %v\n", poped, s)
			wg.Done()
		}()
		wg.Wait()
	}
	log.Println("main end")
}
