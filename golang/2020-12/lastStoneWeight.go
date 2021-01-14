package golang

import (
	"sort"
)

// Last Stone Weight
// We have a collection of stones, each stone has a positive integer weight.
// Each turn, we choose the two heaviest stones and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:
// If x == y, both stones are totally destroyed;
// If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
// At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

func LastStoneWeight(stones []int) int {
	var smash func()
	smash = func() {
		n := len(stones)
		if n <= 1 {
			return
		}
		sort.Ints(stones)
		light := stones[n-2]
		heavy := stones[n-1]
		stones = stones[:n-2]
		if light < heavy {
			stones = append(stones, heavy - light)
		}else {
			stones = append(stones, 0)
		}
		smash()
	}
	smash()
	return stones[0]
}