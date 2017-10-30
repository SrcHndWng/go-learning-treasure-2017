package utils

// Fib calculates Fibonacci.
func Fib(n int) int {
	if (n == 0) || (n == 1) {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}
