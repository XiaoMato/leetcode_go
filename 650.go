package main

import "math"

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	result := 0
	for n != 1 {
		if IsPrime(n) {
			result += n
			return result
		}
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				result += i
				n = n / i
				break
			}
		}
	}
	return result
}

func IsPrime(value int) bool {
	if value < 2 {
		return false
	}
	end := int(math.Sqrt(float64(value)))
	for i := 2; i <= end; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	println(minSteps(16))
}
