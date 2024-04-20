package main

import "math"

func findMinMoves(machines []int) int {
	sum := 0
	for _, v := range machines {
		sum += v
	}

	if sum%len(machines) != 0 {
		return -1
	}

	avg := sum / len(machines)

	step := 0

	for i := 0; i < len(machines)-1; i++ {
		s := avg - machines[i]
		step += int(math.Abs(float64(s)))
		machines[i+1] = machines[i+1] - s
	}
	return step
}
