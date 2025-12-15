struct Solution {}

impl Solution {
    pub fn count_partitions(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut pre_sum = vec![0; nums.len() + 1];
        for (i, &num) in nums.iter().enumerate() {
            pre_sum[i+1] = pre_sum[i] + num;
        }
        for i in 1..nums.len() {
            let l = pre_sum[i];
            let r = pre_sum[nums.len()] - pre_sum[i];
            if (r - l) % 2 == 0 {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    Solution::count_partitions(vec![10,10,3,7,6]);
}