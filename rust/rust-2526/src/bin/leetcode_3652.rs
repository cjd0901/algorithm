struct Solution {}

impl Solution {
    pub fn max_profit(prices: Vec<i32>, strategy: Vec<i32>, k: i32) -> i64 {
        let n = prices.len();
        let mut pre_sum: Vec<i64> = vec![0; n+1];
        for (i, &price) in prices.iter().enumerate() {
            pre_sum[i+1] = pre_sum[i] + (price as i64 * strategy[i] as i64);
        }

        let mut ans = pre_sum[n];
        for i in 0..=n-k as usize{
            let mut sum: i64 = pre_sum[i] + pre_sum[n] - pre_sum[i+k as usize];
            for j in i+(k as usize/2)..i+k as usize {
                sum += prices[j] as i64;
            }
            ans = ans.max(sum);
        }
        ans
    }
}

fn main() {
    let ans = Solution::max_profit(vec![4,7,13], vec![-1,-1,0], 2);
    println!("{}", ans);
}