// Longest Palindromic Substring
// leetcode: https://leetcode-cn.com/problems/longest-palindromic-substring/
struct Solution {}
impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let chs: Vec<char> = s.chars().collect();
        let length = chs.len();
        if length < 2 {
            return s;
        }
        let mut dp = vec![vec![false; length]; length];
        for i in 0..length {
            dp[i][i] = true;
        }
        let (mut max_length, mut begin) = (1, 0);
        for l in 2..length+1 {
            for i in 0..length {
                let j = l + i - 1;
                if j >= length {
                    continue;
                }
                if chs[j] != chs[i] {
                    dp[i][j] = false;
                } else {
                    if l < 3 {
                        dp[i][j] = true;
                    } else {
                        dp[i][j] = dp[i+1][j-1];
                    }
                }
                if dp[i][j] && l > max_length {
                    max_length = l;
                    begin = i;
                }
            }
        }
        chs[begin..begin+max_length].iter().collect()
    }
}

fn main() {
    let ans = Solution::longest_palindrome(String::from("ba"));
    println!("{}", ans);
}