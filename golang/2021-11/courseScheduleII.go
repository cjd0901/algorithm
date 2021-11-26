package _021_11

// Course Schedule II
// leetcode: https://leetcode-cn.com/problems/course-schedule-ii/
func CourseScheduleIIBFS(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		inDegrees = make([]int, numCourses)
		queue = make([]int, 0)
		ans = make([]int, 0)
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		inDegrees[p[0]]++
	}
	for n, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, n)
		}
	}
	for len(queue) > 0 {
		n := queue[0]
		ans = append(ans, n)
		queue = queue[1:]
		for _, edge := range edges[n] {
			inDegrees[edge]--
			if inDegrees[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}

func CourseScheduleIIDFS(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		ans []int
		valid = true
		dfs func(n int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		ans = append(ans, u)
	}

	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}
	for i := 0; i < numCourses/2; i++ {
		ans[i], ans[numCourses-i-1] = ans[numCourses-i-1], ans[i]
	}
	return ans
}