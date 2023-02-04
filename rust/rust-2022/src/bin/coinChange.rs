// Coin Change
// leetcode: https://leetcode-cn.com/problems/coin-change/
struct Solution {}

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let max = (amount + 1) as usize;
        let mut dp = vec![max as i32; max];
        dp[0] = 0;
        for i in 1..max {
            for j in 0..coins.len() {
                if coins[j] <= i as i32 {
                    dp[i] = std::cmp::min(dp[i], dp[i - coins[j] as usize] + 1);
                }
            }
        }
        if dp[amount as usize] > amount { -1 } else { dp[amount as usize] }
    }
}

fn main() {
    let ans = Solution::coin_change(vec![1, 2, 5], 11);
    println!("{}", ans);
}