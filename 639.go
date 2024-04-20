package main

func numDecodings(s string) int {
	if len(s) == 0 {
		return 1
	}
	if len(s) == 1 {
		return one(s)
	}
	o, t := one(s[:1]), two(s[:2])
	if o == 0 && t == 0 {
		return 0
	}
	if o == 0 && t != 0 {
		return t * numDecodings(s[2:]) % (1000000007)
	}
	if o != 0 && t == 0 {
		return o * numDecodings(s[1:]) % (1000000007)
	}
	return (o*numDecodings(s[1:]) + t*numDecodings(s[2:])) % (1000000007)
}

func one(s string) int {
	if s == "*" {
		return 9
	}
	if s == "0" {
		return 0
	}
	return 1
}

func two(s string) int {
	if s[0] == '0' {
		return 0
	}
	if s == "**" {
		return 15
	}
	if s[0] == '*' {
		if s[1] >= '7' {
			return 1
		}
		return 2
	}
	if s[1] == '*' {
		if s[0] == '1' {
			return 9
		}
		if s[0] == '2' {
			return 6
		}
		return 0
	}
	if s > "26" {
		return 0
	}
	return 1
}
