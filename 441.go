package main

func arrangeCoins(n int) int {
	coins := 1
	var result int
	for {
		n = n - coins
		if n < 0 {
			return result
		}
		result++
		coins++
	}
	return result
}
