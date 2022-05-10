package _022_04

import (
	"fmt"
	"math/rand"
	"time"
)

// leetcode: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func FindKthLargest(nums []int, k int) int {
	defer func() {
		fmt.Println(nums)
	}()
	rand.Seed(time.Now().UnixNano())
	return fastSelect(nums, 0, len(nums)-1, k-1)
}

func fastSelect(nums []int, l, r, idx int) int {
	q := randPartition(nums, l, r)
	if q == idx {
		return nums[q]
	} else if q < idx {
		return fastSelect(nums, q+1, r, idx)
	}
	return fastSelect(nums, l, q-1, idx)
}

func randPartition(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[l] = nums[l], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	i := l
	for j := r; j > i; {
		if nums[j] > v {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			j--
		}
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
