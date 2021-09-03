package main

import (
	"fmt"
	"math"
	"sort"
)

// func smallestK(arr []int, k int) []int {
// 	result := make([]int, k+1)
// 	for i, _ := range result {
// 		result[i] = math.MaxInt32
// 	}

// 	for _, v := range arr {
// 		result[0] = v
// 		for i := 1; i < len(result) && result[i] > result[i-1]; i++ {
// 			temp := result[i]
// 			result[i] = result[i-1]
// 			result[i-1] = temp
// 		}
// 	}
// 	return result[1:]
// }

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
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}
