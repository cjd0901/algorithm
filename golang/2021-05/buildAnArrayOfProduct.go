package _021_05

func BuildAnArrayOfProduct(a []int) []int {
	n := len(a)
	b := make([]int, n)
	bottomTriangle, topTriangle := 1, 1
	for i := range b {
		b[i] = 1
	}
	for i := 0; i < n; i++ {
		b[i] *= bottomTriangle
		bottomTriangle *= a[i]

		b[n-i-1] *= topTriangle
		topTriangle *= a[n-1-i]
	}
	return b
}