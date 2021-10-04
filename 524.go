package main

func findLongestWord(s string, dictionary []string) string {
	max := 0
	var result string
	for _, w := range dictionary {
		if len(w) > len(s) {
			continue
		}
		if w == s {
			return s
		}
		index := 0
		for _, c := range s {
			if index >= len(w) {
				break
			}
			if byte(c) == w[index] {
				index++
			}
		}
		if index >= len(w) && len(w) >= max {
			max = len(w)
			if len(w) > len(result) || w < result {
				result = w
			}
		}
	}
	return result
}
