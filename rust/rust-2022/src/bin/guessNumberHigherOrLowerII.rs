// Guess Number Higher or Lower II
// leetcode: https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/
struct Solution {}

impl Solution {
    pub fn get_money_amount(n: i32) -> i32 {
        let len = n as usize;
        let mut dp = vec![vec![0; len+1]; len+1];
        for i in (1..n).rev() {
            for j in i+1..n+1 {
                let mut min_cost = std::i32::MAX;
                let (i, j) = (i as usize, j as usize);
                for k in i..j {
                    min_cost = min_cost.min(k as i32 + std::cmp::max(dp[i][k-1], dp[k+1][j]));
                }
                dp[i][j] = min_cost;
            }
        }
        dp[1][len]
    }
}

fn main() {
    let ans = Solution::get_money_amount(10);
    println!("{}", ans);
}