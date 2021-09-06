package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	if index >= len(nums) || nums[index] != target {
		return -1
	}
	return index
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
