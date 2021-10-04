package main

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	words := strings.Split(text, " ")
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	result := strings.ToLower(strings.Join(words, " "))
	return strings.ToUpper(result[:1]) + result[1:]
}
