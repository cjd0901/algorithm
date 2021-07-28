package _021_07

// Greatest Common Divisor of Strings
// leetcode: https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
func GreatestCommonDivisorOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	for i := min(len1, len2); i >= 1; i-- {
		if len1 % i == 0 && len2 % i == 0 {
			subStr := str2[:i]
			if check(str1, subStr) && check(str2, subStr) {
				return subStr
			}
		}
	}
	return ""
}

func check(str, subStr string) bool {
	n := len(str) / len(subStr)
	tempS := ""
	for i := 0; i < n; i++ {
		tempS += subStr
	}
	return tempS == str
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}