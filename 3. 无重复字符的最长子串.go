package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var chs = []byte(s)
	var n = len(s)
	var freqs = make([]byte, 128)
	var max int
	var left, right int
	for right < n {
		if freqs[chs[right]] == 0 {
			freqs[chs[right]]++
			right++
			max = maxf(max, right - left)
		}else {
			freqs[chs[left]]--
			left++
		}
	}
	return max
}

func maxf(a, b int) int {
	if a > b {
		return a
	}
	return b
}