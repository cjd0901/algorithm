// Longest Increasing Subsequence
// leetcode: https://leetcode-cn.com/problems/longest-increasing-subsequence/
struct Solution {}

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let len = nums.len();
        let mut dp = vec![0; len];
        let mut ans = 0;
        dp[0] = 1;
        for i in 0..len {
            dp[i] = 1;
            for j in 0..i {
                if nums[i] > nums[j] {
                    dp[i] = dp[i].max(dp[j]+1);
                }
            }
            ans = ans.max(dp[i]);
        }
        ans
    }
}

fn main() {
    let ans = Solution::length_of_lis(vec![10,9,2,5,3,7,101,18]);
    println!("{}", ans);
}