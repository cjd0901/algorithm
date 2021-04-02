package _021_03

// Volume of Histogram LCCI
// leetCode: https://leetcode-cn.com/problems/volume-of-histogram-lcci/
// Imagine a histogram (bar graph).
// Design an algorithm to compute the volume of water it could hold if someone poured water across the top.
// You can assume that each histogram bar has width 1.

func VolumeOfHistogramLCCI(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	left := 0
	right := l - 1
	leftH := height[left]
	rightH := height[right]
	ans := 0
	for left < right {
		if leftH < rightH {
			left ++
			if leftH > height[left] {
				ans += leftH - height[left]
			}else {
				leftH = height[left]
			}
		}else {
			right --
			if rightH > height[right] {
				ans += rightH - height[right]
			}else {
				rightH = height[right]
			}
		}
	}
	return ans
}