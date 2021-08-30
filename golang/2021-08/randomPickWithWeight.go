package _021_08

import "math/rand"

// Random Pick with Weight
// https://leetcode-cn.com/problems/random-pick-with-weight/

type Solution struct {
	pre []int
}


func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{pre: w}
}


func (this *Solution) PickIndex() int {
	pick := rand.Intn(this.pre[len(this.pre)-1]) + 1
	l, r := 0, len(this.pre) - 1
	for l < r {
		mid := l + (r-l) >> 2
		if this.pre[mid] < pick {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
