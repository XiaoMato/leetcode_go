package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	int2hex := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
		10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
	}
	result := ""
	time := 0
	for num != 0 {
		remain := num & 15
		result = int2hex[remain] + result
		num = num >> 4
		time++
		if time == 8 {
			break
		}
	}
	var i int
	for i = range result {
		if result[i] == '0' {
			i++
		} else {
			break
		}
	}
	return result[i:]
}