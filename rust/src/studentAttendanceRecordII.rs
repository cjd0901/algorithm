// Student Attendance Record II
// leetcode: https://leetcode-cn.com/problems/student-attendance-record-ii/
struct Solution {}

impl Solution {
    pub fn check_record(n: i32) -> i32 {
        const MOD: i32 = 1e9 as i32+7;
        let mut dp = vec![vec![vec![0; 3]; 2]; (n+1) as usize];
        let mut ans = 0;
        dp[0][0][0] = 1;
        for i in 1..n+1 {
            // 当当前的记录为'P'时
            let i = i as usize;
            for j in 0..2 {
                let j = j as usize;
                for k in 0..3 {
                    let k = k as usize;
                    dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % MOD;
                }
            }
            // 当当前的记录为'A'时
            for k in 0..3 {
                let k = k as usize;
                dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % MOD;
            }
            // 当当前的记录为'L'时
            for j in 0..2 {
                let j = j as usize;
                for k in 1..3 {
                    let k = k as usize;
                    dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % MOD;
                }
            }
        }
        for j in 0..2 {
            for k in 0..3 {
                ans = (dp[n as usize][j as usize][k as usize] + ans) % MOD;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::check_record(8);
    println!("{}", ans);
}