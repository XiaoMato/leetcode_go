package main

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	isPosi := 1
	switch {
	case dividend > 0 && divisor > 0:
		isPosi = 1
	case dividend > 0 && divisor < 0:
		isPosi = -1
		divisor = -divisor
	case dividend < 0 && divisor > 0:
		isPosi = -1
		dividend = -dividend
	case dividend < 0 && divisor < 0:
		isPosi = 1
		divisor = -divisor
		dividend = -dividend
	}
	result := -1
	for dividend > 0 {
		result++
		dividend -= divisor
	}
	return result * isPosi
}
