package main

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, v := range nums {
		var d = v
		if d < 0 {
			d = -d
		}
		if nums[d-1] < 0 {
			result = append(result, d)
			continue
		}
		nums[d-1] = -nums[d-1]
	}
	return result
}
