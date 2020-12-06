package golang

// Pascal's Triangle(杨辉三角)
// leetCode: https://leetcode-cn.com/problems/pascals-triangle/

func PascalTriangle(numRows int) [][]int {
	ans := [][]int{}
	for i := 0; i < numRows; i++ {
		array := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				array = append(array, 1)
			}else {
				array = append(array, ans[i-1][j-1] + ans[i-1][j])
			}
		}
		ans = append(ans, array)
	}
	return ans
}