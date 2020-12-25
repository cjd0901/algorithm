package golang

import "sort"

// Assign Cookies
// leetCode: https://leetcode-cn.com/problems/assign-cookies/
// Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.
// Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j].
// If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content.
// Your goal is to maximize the number of your content children and output the maximum number.

func AssignCookies(g []int, s []int) (maxNumber int) {
	sort.Ints(g)
	sort.Ints(s)
	lg, ls := len(g), len(s)
	for i, j := 0, 0; i < lg && j < ls; i++{
		for j < ls && g[i] > s[j] {
			j++
		}
		if j < ls {
			maxNumber++
			j++
		}
	}
	return
}