// Best Time to Buy and Sell Stock with Cooldown
// leetcode: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
struct Solution {}
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let len = prices.len();
        // dp[i][0] 代表当日持有股票   dp[i][1] 代表当日未持有股票并在冷冻期  dp[i][2] 代表当日未持有股票并不在冷冻期
        let mut dp = vec![vec![0; 3]; len];
        dp[0][0] = -prices[0];
        for i in 1..len {
            dp[i][0] = (dp[i-1][2] - prices[i]).max(dp[i-1][0]);
            dp[i][1] = dp[i-1][0] + prices[i];
            dp[i][2] = dp[i-1][1].max(dp[i-1][2]);
        }
        dp[len-1][0].max(dp[len-1][1]).max(dp[len-1][2])
    }

    pub fn max_profit_simple(prices: Vec<i32>) -> i32 {
        let len = prices.len();
        // q 代表当日持有股票   w 代表当日未持有股票并在冷冻期  e 代表当日未持有股票并不在冷冻期
        let (mut q, mut w, mut e) = (-prices[0], 0, 0);
        for i in 1..len {
            let (q1, w1, e1) = (q, w, e);
            q = q1.max(e1 - prices[i]);
            w = q1 + prices[i];
            e = w1.max(e1);
        }
        q.max(w).max(e)
    }
}

fn main() {
    let ans = Solution::max_profit_simple(vec![1,2,3,0,2]);
    println!("{:?}", ans);
}