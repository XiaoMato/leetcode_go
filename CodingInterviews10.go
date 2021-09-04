package main

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first, second, third := 0, 1, 1

	for i := 2; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}

	return second
}
