package golang

// Split Array into Consecutive Subsequences
// leetCode: https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/

func SplitArray(nums []int) bool {
	l := len(nums)
	isPossible := true
	if l < 3 {
		return false
	}
	max := []int{}
	size := []int{}
	for index, n := range nums {
		full := false
		func(full *bool) {
			for i := len(size) - 1; i >= 0; i-- {
				if n == max[i] + 1 {
					max[i] = n
					size[i]++
					*full = false
					break
				}else {
					*full = true
				}
			}
		}(&full)
		if index == 0 || full {
			size = append(size, 1)
			max = append(max, n)
		}
		//fmt.Println(full)
	}
	//fmt.Println("size", size)
	//fmt.Println("max", max)
	for _, s := range size {
		if s < 3 {
			isPossible = false
			break
		}
	}
	return isPossible
}