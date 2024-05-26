package main

func getWinner(arr []int, k int) int {
	maxInt := arr[0]
	for _, v := range arr {
		if v > maxInt {
			maxInt = v
		}
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var res int
	var lastWin = max(arr[0], arr[1])
	for res < k {
		if arr[0] == maxInt || arr[1] == maxInt {
			return maxInt
		}
		if arr[0] > arr[1] {
			if arr[0] == lastWin {
				res++
			} else {
				lastWin = arr[0]
				res = 1
			}
			arr = append(arr[:1], arr[2:]...)
		} else {
			if arr[1] == lastWin {
				res++
			} else {
				lastWin = arr[1]
				res = 1
			}
			arr = append(arr[1:], arr[0])
		}
	}
	return lastWin
}
