package utils

import (
	"strconv"
	"testing"
)

func TestFib(t *testing.T) {
	test := func(arg int, exp int) {
		if Fib(arg) != exp {
			t.Error("Fib returns error. arg = " + strconv.Itoa(arg))
		}
	}

	test(0, 0)
	test(1, 1)
	test(2, 1)
	test(3, 2)
	test(5, 5)
	test(10, 55)
	test(20, 6765)
}
