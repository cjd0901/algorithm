package _021_05

// Number of Submatrices That Sum to Target
// leetcode: https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/
// Given a matrix and a target, return the number of non-empty submatrices that sum to target.
// A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
// Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.

func NumberOfSubmatricesThatSumToTarget(matrix [][]int, target int) int {
	count := 0
	for i := range matrix {
		sumList := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for j, n := range row {
				sumList[j] += n
			}
			count += subArrayEqualsK(sumList, target)
		}
	}
	return count
}

func subArrayEqualsK(nums []int, k int) int {
	count, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre]++
	}
	return count
}