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

func main() {
	println(numDecodings("1*6*7*1*9*6*2*9*2*3*3*6*3*2*2*4*7*2*9*6*0*6*4*4*1*6*9*0*5*9*2*5*7*7*0*6*9*7*1*5*5*9*3*0*4*9*2*6*2*5*7*6*1*9*4*5*8*4*7*4*2*7*1*2*1*9*1*3*0*6*"))
	//println(numDecodings("1*"))
}
