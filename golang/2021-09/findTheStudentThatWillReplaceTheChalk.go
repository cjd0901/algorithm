package _021_09

// Find the Student that Will Replace the Chalk
// https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk/
func FindTheStudentThatWillReplaceTheChalk(chalk []int, k int) int {
	sum := 0
	for _, c := range chalk {
		sum += c
	}
	remainder := k%sum
	for i, c := range chalk {
		remainder -= c
		if remainder < 0 {
			return i
		}
	}
	return -1
}

func FindTheStudentThatWillReplaceTheChalk2(chalk []int, k int) int {
	for i := 1; i < len(chalk); i++ {
		chalk[i] += chalk[i-1]
	}
	k = k%chalk[len(chalk)-1]
	l, r := 0, len(chalk)-1
	for k < r {
		mid := l + (r -l) >> 1
		if chalk[mid] <= k {
			l++
		} else {
			r--
		}
	}
	return l
}