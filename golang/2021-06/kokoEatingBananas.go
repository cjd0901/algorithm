package _021_06

// Koko Eating Bananas
// leetcode: https://leetcode-cn.com/problems/koko-eating-bananas/

func KokoEatingBananas(piles []int, h int) int {
	lo := 1
	hi := 0
	for _, pile := range piles {
		if pile > hi {
			hi = pile
		}
	}
	for lo < hi {
		mid := (lo + hi) / 2
		if !possible(h, mid, piles) {
			lo = mid + 1
		}else {
			hi = mid
		}
	}
	return lo
}

func possible(h, c int, piles []int) bool {
	t := 0
	for _, pile := range piles {
		t += (pile-1)/c + 1
	}
	return t <= h
}