package stack

import (
	"fmt"
	"strconv"
	"testing"
)

var s Stack

const limit = 5

var ary = []string{"a", "b", "c", "d", "e", "f", "g"}

func TestPush(t *testing.T) {
	before()
	printStack()

	if len(s.vals) != limit {
		t.Error("Push length error.")
	}
	if s.vals[0] != ary[len(ary)-limit] {
		t.Error("Push stack value error.")
	}
	if s.vals[limit-1] != ary[len(ary)-1] {
		t.Error("Push stack value error.")
	}
}

func TestPop(t *testing.T) {
	test := func(result string, expect string, lenght int) {
		printStack()
		if result != expect {
			t.Error("Pop value error. expect = " + expect)
		}
		if len(s.vals) != lenght {
			t.Error("Pop length error. expect length = " + strconv.Itoa(lenght))
		}
	}

	before()

	for i := len(ary) - 1; i >= len(ary)-limit; i-- {
		test(s.Pop(), ary[i], i-(len(ary)-limit))
	}
}

func before() {
	s = NewStack(limit)
	for _, v := range ary {
		s.Push(v)
	}
}

func printStack() {
	fmt.Printf("Stack vals = %v\n", s.vals)
}
