package _021_12

// Range Sum Query - Mutable
// leetcode: https://leetcode-cn.com/problems/range-sum-query-mutable/
type NumArray struct {
	_nums, old []int
}

func lowBit(x int) int {
	return x&(-x)
}

func (this *NumArray) Add(i, v int) {
	for j := i; j < len(this._nums); j += lowBit(j) {
		this._nums[j] += v
	}
}

func (this *NumArray) Query(i int) int {
	sum := 0
	for j := i; j > 0; j -= lowBit(j) {
		sum += this._nums[j]
	}
	return sum
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	_nums := make([]int, l+1)
	numArray := NumArray{_nums: _nums, old: nums}
	for i := 0; i < l; i++ {
		numArray.Add(i+1, nums[i])
	}
	return numArray
}


func (this *NumArray) Update(index int, val int)  {
	this.Add(index+1, val - this.old[index])
	this.old[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(right+1) - this.Query(left)
}
