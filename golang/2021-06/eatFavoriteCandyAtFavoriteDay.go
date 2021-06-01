package _021_06

// Can You Eat Your Favorite Candy on Your Favorite Day?
// leetcode: https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/

func EatFavoriteCandyAtFavoriteDay(candiesCount []int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	candiesSum := make([]int, len(candiesCount))
	candiesSum[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++ {
		candiesSum[i] = candiesCount[i] + candiesSum[i-1]
	}
	for i, query := range queries {
		ft, fd, cd := query[0], query[1], query[2]
		x1 := fd + 1
		y1 := (fd + 1) * cd
		x2 := 1
		if ft > 0 {
			x2 = candiesSum[ft-1] + 1
		}
		y2 := candiesSum[ft]
		ans[i] = !(x1 > y2 || y1 < x2)
	}
	return ans
}