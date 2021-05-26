package _021_05

// Reverse Substrings Between Each Pair of Parentheses
// leetcode: https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/
// You are given a string s that consists of lower case English letters and brackets. 
// Reverse the strings in each pair of matching parentheses, starting from the innermost one.
// Your result should not contain any brackets.

func ReverseSubstringsBetweenEachPairOfParentheses(s string) string {
	stack, str := [][]byte{}, []byte{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		}else if s[i] == ')' {
			n := len(str)
			for j := 0; j < n / 2; j++ {
				str[j], str[n-j-1] = str[n-j-1], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		}else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
