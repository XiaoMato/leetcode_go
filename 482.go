package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	l := len(s)
	r := l % k
	result := s[:r]
	remain := s[r:]
	for remain != "" {
		if len(result) > 0 {
			result = result + "-"
		}
		result = result + remain[:k]
		remain = remain[k:]
	}
	return strings.ToUpper(result)
}
