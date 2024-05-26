package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if abs(i-j) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
