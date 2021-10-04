package main

func balancedStringSplit(s string) int {
	var r, l, result int
	for _, v := range s {
		if v == 'R' {
			r++
		} else {
			l++
		}

		if r == l {
			result++
			r, l = 0, 0
		}
	}
	return result
}
