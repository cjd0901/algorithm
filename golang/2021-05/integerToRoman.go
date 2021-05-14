package _021_05

// Integer to Roman
// leetcode: https://leetcode-cn.com/problems/integer-to-roman/
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

func IntegerToRoman(num int) string {
	i2r := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	romans := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ans := ""
	for i := 0; i < len(romans) && num > 0; i++ {
		d := num / romans[i]
		if d < 1 {
			continue
		}
		for j := 0; j < d; j++ {
			ans += i2r[romans[i]]
		}
		num = num % romans[i]
	}
	return ans
}
