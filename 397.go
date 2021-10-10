package main

import "math/bits"

func integerReplacement(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	if n&1 == 0 {
		return 1 + integerReplacement(n/2)
	}
	if bits.OnesCount(uint(n+1)) == 1 {
		return 1 + integerReplacement(n+1)
	}
	if bits.OnesCount(uint(n-1)) == 1 {
		return 1 + integerReplacement(n-1)
	}
	return 1 + min(integerReplacement(n+1), integerReplacement(n-1))

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(integerReplacement(13))
}
