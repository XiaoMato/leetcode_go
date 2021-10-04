package main

func checkValidString(s string) bool {
	return ok(s)
}

func ok(s string) bool {
	var star, left int
	for _, v := range s {
		switch v {
		case '(':
			left++
		case '*':
			star++
		case ')':
			if star == 0 && left == 0 {
				return false
			}
			if star > 0 && left == 0 {
				star--
			} else if left > 0 {
				left--
			}
		}
	}
	if left > 0 {
		return false
	}
	return true
}

func reverse(s string) bool {
	var star, right int
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ')':
			right++
		case '*':
			star++
		case '(':
			if star == 0 && right == 0 {
				return false
			}
			if star > 0 && right == 0 {
				star--
			} else if right > 0 {
				right--
			}
		}
	}
	if right > 0 {
		return false
	}
	return true
}
