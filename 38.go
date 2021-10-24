package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	result := "1"
	for i := 0; i < n-1; i++ {
		tmp := ""
		count := 1
		for k := 1; k < len(result); k++ {
			if result[k] == result[k-1] {
				count++
			} else {
				tmp = fmt.Sprintf("%v%v%c", tmp, count, result[k-1])
				count = 1
			}
		}
		tmp = fmt.Sprintf("%v%v%c", tmp, count, result[len(result)-1])
		result = tmp
	}
	return fmt.Sprint(result)
}

func main() {
	println(countAndSay(4))
}
