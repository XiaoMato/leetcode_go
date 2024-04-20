package main 

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if m[target - v] != 0 && m[target-v] != i {
			return []int{i, m[target-v]}
		}
	}
	return nil
}