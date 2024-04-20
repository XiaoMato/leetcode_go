package main

import "math"

func findNumberOfLIS(nums []int) int {
	mark := make([]int, len(nums))
	for i := range mark {
		mark[i] = 1
	}

	var max = -1
	var smax = -1

	for i := 0; i < len(mark); i++ {
		for j := i + 1; j < len(mark); j++ {
			if nums[i] < nums[j] {
				mark[j] = int(math.Max(float64(mark[j]), float64(mark[i]+1)))
				//mark[j]++
				if mark[j] > max {
					smax = max
					max = mark[j]
				}
			}
		}
	}

	if smax == -1 {
		return len(nums)
	}

	result := 0
	for _, v := range mark {
		if v == smax {
			result++
		}
	}
	return result
}