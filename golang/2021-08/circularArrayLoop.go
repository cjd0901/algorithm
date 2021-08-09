package _021_08

// Circular Array Loop
// leetcode: https://leetcode-cn.com/problems/circular-array-loop/
func CircularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(i int) int {
		return ((nums[i] + i) % n + n) % n
	}
	for i, num := range nums {
		if num == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		a := i
		for nums[a]*nums[next(a)] > 0 {
			t := a
			a = next(a)
			nums[t] = 0
		}
	}
	return false
}