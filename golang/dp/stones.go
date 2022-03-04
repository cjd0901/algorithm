package dp

// Stones
// https://atcoder.jp/contests/dp/tasks/dp_k
func Stones(K int, nums []int) string {
	win := make([]bool, K+1)
	for k := 1; k <= K; k++ {
		for _, num := range nums {
			if k >= num && !win[k-num] {
				win[k] = true
				break
			}
		}
	}
	ans := "Second"
	if win[K] {
		ans = "First"
	}
	return ans
}