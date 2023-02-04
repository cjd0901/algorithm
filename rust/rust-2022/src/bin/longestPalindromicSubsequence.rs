// Longest Palindromic Subsequence
// leetcode: https://leetcode-cn.com/problems/longest-palindromic-subsequence/
struct Solution {}
impl Solution {
    pub fn longest_palindrome_subseq(s: String) -> i32 {
        let n = s.len();
        let b = s.as_bytes();
        let mut dp = vec![vec![0; n]; n];
        for i in (0..n).rev() {
            dp[i][i] = 1;
            for j in i+1..n {
                if b[i] == b[j] {
                    dp[i][j] = dp[i+1][j-1]+2;
                } else {
                    dp[i][j] = dp[i+1][j].max(dp[i][j-1]);
                }
            }
        }
        dp[0][n-1]
    }
}

fn main() {
    let ans = Solution::longest_palindrome_subseq(String::from("bbbbab"));
    println!("{}", ans);
}