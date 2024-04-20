package main

import "slices"


func longestPalindrome(s string) string {
	bytes := []byte(s)
	reverseBytes := slices.Clone(bytes)
	slices.Reverse(reverseBytes)
	n := len(s)
	var m [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if reverseBytes[0] == bytes[i] {
			m[0][i] = 1
		}
		if bytes[0] == reverseBytes[i] {
			m[i][0] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if bytes[j] == reverseBytes[i] {
				m[i][j] = m[i-1][j-1] + 1
			} else  {
				m[i][j] = 0
			}
		}
	}
	max := m[n-1][n-1]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == max {
				return s[j-max+1:j+1]
			}
		}
	}
	return ""
}