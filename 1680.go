package main

import (
	"math"
	"strconv"
)

func concatenatedBinary(n int) int {
	primeBase := int(math.Pow(float64(10), float64(9))) + 7
	result := 0
	remains := 1
	move := 0
	two := 1

	for i := n; i >= 1; i-- {
		two = int((uint64(math.Pow(float64(2), float64(move)))%uint64(primeBase))*(uint64(two))) % primeBase
		remains = int((uint64(i) * uint64(two)) % uint64(primeBase))
		result = (result + remains) % primeBase
		move = len(strconv.FormatInt(int64(i), 2))
	}
	return result
}
