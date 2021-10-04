package main

import "sort"

type pair struct {
	strong int
	array  int
}

func kWeakestRows(mat [][]int, k int) []int {
	pairs := make([]pair, 0)
	for i, v := range mat {
		count := 0
		for _, t := range v {
			if t == 1 {
				count++
			}
		}
		pairs = append(pairs, pair{strong: count, array: i})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].strong < pairs[j].strong
	})
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].array)
	}
	return result
}
