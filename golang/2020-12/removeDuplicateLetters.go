package golang

// Remove Duplicate Letters
// leetCode: https://leetcode-cn.com/problems/remove-duplicate-letters/
// Given a string s, remove duplicate letters so that every letter appears once and only once.
// You must make sure your result is the smallest in lexicographical order among all possible results.

// 栈 num->每个字母的个数, stack->栈, inStack->是否在栈中
func RemoveDuplicateLetters(s string) string {
	num := make([]int, 26)
	for _, ch := range s {
		num[ch - 'a']++
	}
	stack := make([]byte, 0)
	inStack := make([]bool, 26)
	for i := range s {
		ch := s[i]
		if !inStack[ch - 'a'] {
			for len(stack) > 0 && ch < stack[len(stack) - 1] {
				last := stack[len(stack) - 1] - 'a'
				if num[last] == 0 {
					break
				}
				stack = stack[:len(stack) - 1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch - 'a'] = true
		}
		num[ch - 'a']--
	}
	return string(stack)
}