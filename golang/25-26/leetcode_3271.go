package golang

func stringHash(s string, k int) string {
	bytes := []byte(s)
	n := len(s)
	result := make([]byte, n/k)
	for i := 0; i < n; i += k {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += int(bytes[j] - 'a')
		}
		result[i/k] = byte(sum%26) + 'a'
	}
	return string(result)
}
