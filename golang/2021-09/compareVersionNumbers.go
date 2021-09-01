package _021_09

import (
	"strconv"
	"strings"
)

// Compare Version Numbers
// leetcode: https://leetcode-cn.com/problems/compare-version-numbers/

func CompareVersionNumbers(version1 string, version2 string) int {
	v1Part := strings.Split(version1, ".")
	v2Part := strings.Split(version2, ".")
	for i := 0; i < len(v1Part) || i < len(v2Part); i++ {
		x, y := 0, 0
		if i < len(v1Part) {
			x, _ = strconv.Atoi(v1Part[i])
		}
		if i < len(v2Part) {
			y, _ = strconv.Atoi(v2Part[i])
		}
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}

func CompareVersionNumbers2(version1, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0
	for i < m || j < n {
		x := 0
		for ; i < m && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i] - '0')
		}
		i++
		y := 0
		for ; j < n && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j] - '0')
		}
		j++
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}