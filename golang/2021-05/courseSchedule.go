package _021_05

// Course Schedule
// leetcode: https://leetcode-cn.com/problems/course-schedule/
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.

// TODO: 拓扑排列 & DFS & BFS
func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	hasTook := make(map[int]bool)

}