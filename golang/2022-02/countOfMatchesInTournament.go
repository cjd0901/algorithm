package _022_02

// Count of Matches in Tournament
// leetcode: https://leetcode-cn.com/problems/count-of-matches-in-tournament/
func CountOfMatchesInTournament(n int) int {
	match := 0
	for n > 1 {
		match += n/2
		n = n/2 + n%2
	}
	return match
}