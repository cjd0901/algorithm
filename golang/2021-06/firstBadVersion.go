package _021_06

func FirstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := (lo + hi) / 2
		if !isBadVersion(mid) {
			lo = mid + 1
		}else {
			hi = mid
		}
	}
	return lo
}