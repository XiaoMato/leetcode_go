package main

func lengthOfLastWord(s string) int {
	result := 0

	index := len(s) - 1
	for ; index >= 0; index-- {
		if s[index] != ' ' {
			break
		}
	}

	for i := index; i >= 0; i-- {
		if s[i] != ' ' {
			result++
		} else {
			return result
		}
	}
	return result
}

func main() {
	println(lengthOfLastWord("luffy is still joyboy  "))
}
