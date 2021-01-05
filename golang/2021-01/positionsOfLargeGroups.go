package _021_01

// Positions Of Large Groups
// leetCode: https://leetcode-cn.com/problems/positions-of-large-groups/
// In a string s of lowercase letters, these letters form consecutive groups of the same character.
// For example, a string like s = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z", and "yy".
// A group is identified by an interval [start, end], where start and end denote the start and end indices (inclusive) of the group. In the above example, "xxxx" has the interval [3,6].
// A group is considered large if it has 3 or more characters.
// Return the intervals of every large group sorted in increasing order by start index.

func LargeGroupPositions(s string) [][]int {
	var start, prev int
	ans := make([][]int, 0)
	for i, ch := range s {
		if int(ch) != prev {
			if i-start >= 3 {
				group := []int{start, i-1}
				ans = append(ans, group)
			}
			prev = int(ch)
			start = i
		}
	}
	if int(s[len(s)-1]) == prev {
		if len(s)-1-start+1 >= 3 {
			group := []int{start, len(s)-1}
			ans = append(ans, group)
		}
	}
	return ans
}

func LargeGroupPositions2(s string) [][]int {
	ans := make([][]int, 0)
	cnt := 1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				ans = append(ans, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return ans
}