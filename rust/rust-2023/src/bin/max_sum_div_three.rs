struct Solution {}

use std::cmp::max;

impl Solution {
    pub fn max_sum_div_three(nums: Vec<i32>) -> i32 {
        let len = nums.len();
        let mut dp = vec![vec![0; 3]; len];
        for i in 0..len {
            let num = nums[i];
            let num_mod = num % 3;
            match num_mod {
                0 => {
                    if i == 0 {
                        dp[i][0] = num;
                    } else {
                        dp[i][0] = dp[i - 1][0] + num;
                        if dp[i - 1][1] != 0 {
                            dp[i][1] = dp[i - 1][1] + num;
                        }
                        if dp[i - 1][2] != 0 {
                            dp[i][2] = dp[i - 1][2] + num;
                        }
                    }
                }
                1 => {
                    if i == 0 {
                        dp[i][1] = num;
                    } else {
                        dp[i][0] = dp[i - 1][0];
                        dp[i][2] = dp[i - 1][2];
                        dp[i][1] = max(dp[i - 1][0] + num, dp[i - 1][1]);
                        if dp[i - 1][2] != 0 {
                            dp[i][0] = max(dp[i - 1][2] + num, dp[i][0]);
                        }
                        if dp[i - 1][1] != 0 {
                            dp[i][2] = max(dp[i - 1][1] + num, dp[i][2]);
                        }
                    }
                }
                _ => {
                    if i == 0 {
                        dp[i][2] = num;
                    } else {
                        dp[i][0] = dp[i - 1][0];
                        dp[i][1] = dp[i - 1][1];
                        dp[i][2] = max(dp[i - 1][0] + num, dp[i - 1][2]);
                        if dp[i - 1][1] != 0 {
                            dp[i][0] = max(dp[i - 1][1] + num, dp[i][0]);
                        }
                        if dp[i - 1][2] != 0 {
                            dp[i][1] = max(dp[i - 1][2] + num, dp[i][1]);
                        }
                    }
                }
            }
        }
        dp[len - 1][0]
    }
}

fn main() {}
