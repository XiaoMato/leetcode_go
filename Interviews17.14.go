package main

import (
	"math"
)

func smallestK(arr []int, k int) []int {
	result := make([]int, k+1)
	for i, _ := range result {
		result[i] = math.MaxInt32
	}

	for _, v := range arr {
		result[0] = v
		for i := 1; i < len(result) && result[i] > result[i-1]; i++ {
			temp := result[i]
			result[i] = result[i-1]
			result[i-1] = temp
		}
	}
	return result[1:]
}
