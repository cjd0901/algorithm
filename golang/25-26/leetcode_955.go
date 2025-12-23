package golang

func minDeletionSize(strs []string) int {
	ans := 0
	n, m := len(strs), len(strs[0])
	cur := make([]string, n)
next:
	for j := range m {
		for i := range n - 1 {
			if cur[i]+string(strs[i][j]) > cur[i+1]+string(strs[i+1][j]) {
				ans++
				continue next
			}
		}
		for i, s := range strs {
			cur[i] += string(s[j])
		}
	}
	return ans
}
