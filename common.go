package main


func minInt(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}