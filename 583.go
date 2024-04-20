package main

import "math"

func minDistance(word1 string, word2 string) int {
	word1, word2 = " "+word1, " "+word2
	x, y := len(word1), len(word2)
	mark := make([][]int, x)
	for i, _ := range mark {
		mark[i] = make([]int, y)
	}
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if word1[i] == word2[j] {
				mark[i][j] = int(math.Max(math.Max(float64(mark[i-1][j]), float64(mark[i][j-1])), float64(mark[i-1][j-1])+1))
			} else {
				mark[i][j] = int(math.Max(math.Max(float64(mark[i-1][j]), float64(mark[i][j-1])), float64(mark[i-1][j-1])))
			}
		}
	}
	return x - y - mark[x-1][y-1] - 2
}
