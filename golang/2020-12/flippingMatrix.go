package golang

import (
	"math"
)

// Score After Flipping Matrix
// leetCode: https://leetcode-cn.com/problems/score-after-flipping-matrix/

func FlippingMatrix(A [][]int) int {
	rowN := len(A)
	cowN := len(A[0])
	nums := make([]int, cowN)
	ans := 0.0
	for i := 0; i < rowN; i++ {
		if A[i][0] == 0 {
			for j := 0; j < cowN; j++ {
				if A[i][j] == 0 {
					nums[j] += 1
				}
			}
		}else {
			for j := 0; j < cowN; j++ {
				if A[i][j] == 1 {
					nums[j] += 1
				}
			}
		}
	}
	for index, n := range nums {
		if n <= rowN / 2 {
			ans += math.Pow(2.0, float64(cowN - index - 1)) * float64(rowN - n)
		}else {
			ans += math.Pow(2.0, float64(cowN - index - 1)) * float64(n)
		}
	}
	return int(ans)
}