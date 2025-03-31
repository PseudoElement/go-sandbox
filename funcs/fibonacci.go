package funcs

func Fibonacci(order int) int {
	if order == 0 {
		return -1
	}
	return fib(order, 1, 0)
}

func fib(order int, curr int, prev int) int {
	order--
	if order == 0 {
		return prev
	}

	return fib(order, curr+prev, curr)
}
