package _021_05

// Course Schedule
// leetcode: https://leetcode-cn.com/problems/course-schedule/
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.

// 拓扑排列 DFS
func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	// searched 0 => 未搜索 1 => 搜索中 2 => 已搜索
	searched := make([]int, numCourses)
	valid := true
	var dfs func(c int)
	dfs = func(c int) {
		searched[c] = 1
		for _, v := range edges[c] {
			if searched[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}else if searched[v] == 1 {
				valid = false
				return
			}
		}
		searched[c] = 2
	}

	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if searched[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// 拓扑排列 BFS
func CourseSchedule2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	result := 0
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indegree[prerequisite[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		result++
		for _, v := range edges[c] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return result == numCourses
}