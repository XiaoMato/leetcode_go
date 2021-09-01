package main

import (
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	if zero(version1) && zero(version2) {
		return 0
	}
	for version1 != "" && version2 != "" {
		prefix1 := getPre(version1)
		prefix2 := getPre(version2)
		int1 := toInt(prefix1)
		int2 := toInt(prefix2)
		if int1 > int2 {
			return 1
		}
		if int1 < int2 {
			return -1
		}
		version1 = trim(version1)
		version2 = trim(version2)
	}
	if zero(version1) && zero(version2) {
		return 0
	}
	if zero(version1) && !zero(version2) {
		return -1
	}
	if !zero(version1) && zero(version2) {
		return 1
	}
	return 0
}

func getPre(version string) string {
	var result []byte
	for _, v := range version {
		if v == '.' {
			return string(result)
		}
		result = append(result, byte(v))
	}
	return string(result)
}

func toInt(version string) int {
	index := 0
	for _, v := range version {
		if v == '0' {
			index++
		} else {
			break
		}
	}
	if index == len(version) {
		return 0
	}
	i, _ := strconv.ParseInt(version[index:], 10, 0)
	return int(i)
}

func zero(version string) bool {
	if version == "" {
		return true
	}
	for _, v := range version {
		if v != '0' && v != '.' {
			return false
		}
	}
	return true
}

func trim(version string) string {
	index := 0
	for _, v := range version {
		if v != '.' {
			index++
			continue
		}
		break
	}
	if index == len(version) {
		return ""
	}
	return version[index+1:]
}
