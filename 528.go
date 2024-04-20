package main

import "math/rand"

type Solution struct {
	prefixSum []int
}

func PickIndexConstructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	prefixSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		prefixSum[i] = prefixSum[i-1] + w[i]
	}
	return Solution{
		prefixSum: prefixSum,
	}
}

func (this *Solution) PickIndex() int {
	i := 0
	n := rand.Intn(this.prefixSum[len(this.prefixSum)-1])
	for _, v := range this.prefixSum {
		if v > n {
			i++
		}
	}
	return i
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
