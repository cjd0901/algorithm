package _022_02

// Sum of Unique Elements
// leetcode: https://leetcode-cn.com/problems/sum-of-unique-elements/
func SumOfUniqueElements(nums []int) int {
	sum := 0
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	for k, v := range countMap {
		if v == 1 {
			sum += k
		}
	}
	return sum
}