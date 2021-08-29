package main

func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)

	result := 0

	for i := 1; i <= length; i += 2 {
		result += subSum(arr, i)
	}
	return result
}

func subSum(arr []int, length int) int {
	sum := preSum(arr, length)
	pre := sum
	for i := length; i < len(arr); i++ {
		pre += arr[i]
		pre -= arr[i-length]
		sum += pre
	}
	return sum
}

func preSum(arr []int, prefix int) int {
	sum := 0
	for i := 0; i < prefix; i++ {
		sum += arr[i]
	}
	return sum
}
