package main

import (
	"math"
)

func integerBreak(n int) int {
	t := int(math.Floor(math.Sqrt(float64(n))))
	sum := (n/t)*t
	nums := make([]int, 0)
	for i := 0; i < t; i++ {
		nums = append(nums, t)
	}
	index := 0
	for sum != n {
		nums[index]++
		sum++
	}
	result := 1

	for _, v := range nums {
		result *= v
	}
	return result
}
