package main

import "fmt"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = fmt.Sprint(i + 1)
	}
	index := 2
	for index < n {
		result[index] = "Fizz"
		index += 3
	}
	index = 4
	for index < n {
		result[index] = "Buzz"
		index += 5
	}
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			result[i] = "FizzBuzz"
		}
	}
	return result
}
