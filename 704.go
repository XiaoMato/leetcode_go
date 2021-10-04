package main

import (
	"sort"
)

func search(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	if index >= len(nums) || nums[index] != target {
		return -1
	}
	return index
}
