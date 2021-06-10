package _021_06

// Course Schedule IV
// leetcode: https://leetcode-cn.com/problems/course-schedule-iv/

func CourseScheduleIV(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	queryRes := make([]bool, len(queries))
	preList := make([][]int, numCourses)
	for i := range preList {
		preList[i] = make([]int, 0)
	}
	var dfs func(c, q int) bool
	dfs = func(c, q int) bool {
		if c == q {
			return true
		}
		for _, p := range preList[c] {
			if dfs(p, q) {
				return true
			}
		}
		return false
	}
	for _, prerequisite := range prerequisites {
		c := prerequisite[1]
		p := prerequisite[0]
		preList[c] = append(preList[c], p)
	}
	for i, query := range queries {
		queryRes[i] = dfs(query[1], query[0])
	}
	return queryRes
}