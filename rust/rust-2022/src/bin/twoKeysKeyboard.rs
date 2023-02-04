// 2 Keys Keyboard
// leetcode: 2 Keys Keyboard
struct Solution {}
impl Solution {
    pub fn min_steps(n: i32) -> i32 {
        let n = n as usize;
        let mut dp = vec![i32::MAX; n+1];
        dp[1] = 0;
        let mut i = 2 as usize;
        while i <= n {
            let mut j = 1 as usize;
            while j * j <= i {
                if i % j == 0 {
                    dp[i] = dp[i].min(dp[j] + (i/j) as i32);
                }
                j += 1;
            }
            i += 1;
        }
        println!("{:?}", dp);
        dp[n] as i32
    }
}

fn main() {
    let ans = Solution::min_steps(3);
    println!("{}", ans);
}