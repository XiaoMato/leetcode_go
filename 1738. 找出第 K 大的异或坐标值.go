package main

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	for i := 1; i < len(matrix); i++ {
		matrix[i][0] = matrix[i-1][0] ^ matrix[i][0]
	}

	for i := 1; i < len(matrix[0]); i++ {
		matrix[0][i] = matrix[0][i-1] ^ matrix[0][i]
	}
	var pre []int
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] = matrix[i-1][j] ^ matrix[i][j-1] ^ matrix[i-1][j-1] ^ matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		pre = append(pre, matrix[i]...)
	}
	sort.Ints(pre)
	return pre[len(pre)-k]
}
