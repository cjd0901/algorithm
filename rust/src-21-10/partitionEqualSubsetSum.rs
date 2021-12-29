// Partition Equal Subset Sum
// leetcode: https://leetcode-cn.com/problems/partition-equal-subset-sum/
struct Solution {}

impl Solution {
    pub fn can_partition(nums: Vec<i32>) -> bool {
        let mut sum = 0;
        let mut max = 0;
        for num in nums.iter() {
            sum += *num;
            if *num > max {
                max = *num;
            }
        }
        if sum % 2 == 1 || sum / 2 < max {
            return false;
        }

        let target = sum / 2;
        let len = nums.len();
        let mut dp = vec![vec![false; (target+1) as usize]; len];
        for i in 0..len {
            dp[i][0] = true;
        }
        dp[0][nums[0] as usize] = true;

        for i in 1..len {
            for j in 1..=target as usize {
                if j as i32 > nums[i] {
                    dp[i][j] = dp[i-1][j - nums[i] as usize] || dp[i-1][j];
                } else {
                    dp[i][j] = dp[i-1][j];
                }
            }
        }
        dp[len-1][target as usize]
    }

    pub fn can_partition2(nums: Vec<i32>) -> bool {
        let len = nums.len();
        if len < 2 {
            return false;
        }
        let (mut max, mut sum) = (0, 0);
        for num in nums.iter() {
            sum += *num;
            if *num > max {
                max = *num;
            }
        }
        if sum % 2 == 1 || sum / 2 < max {
            return false;
        }
        let target = (sum / 2) as usize;
        let mut dp = vec![false; target + 1];
        dp[0] = true;
        for i in 0..len {
            let num = nums[i];
            for j in (num as usize ..= target).rev() {
                dp[j] = dp[j] || dp[j-num as usize];
            }
        }
        dp[target]
    }
}

fn main() {
    let ans = Solution::can_partition2(vec![1,5,11,5]);
    println!("{}", ans);
}