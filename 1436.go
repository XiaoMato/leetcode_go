package main

func destCity(paths [][]string) string {
	var count, p map[string]int
	for _, v := range paths {
		count[v[0]] = count[v[0]] + 1
		count[v[1]] = count[v[1]] + 1
		p[v[0]] = 1
		p[v[1]] = 2
	}

	for k, v := range count {
		if v == 1 && p[k] == 2 {
			return k
		}
	}
	return ""
}
