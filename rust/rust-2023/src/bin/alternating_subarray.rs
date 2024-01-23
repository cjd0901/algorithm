struct Solution {}

impl Solution {
    pub fn alternating_subarray(nums: Vec<i32>) -> i32 {
        let (mut ans, mut start) = (-1, 0);
        let mut valid = false;
        for i in 0..nums.len() - 1 {
            if !valid && nums[i] + 1 == nums[i + 1] {
                valid = true;
                start = i;
                ans = ans.max(2);
            } else if valid {
                if nums[i - 1] == nums[i + 1] {
                    ans = ans.max((i - start) as i32 + 2);
                } else if nums[i] + 1 == nums[i + 1] {
                    start = i;
                } else {
                    valid = false;
                }
            }
        }
        ans
    }
}

fn main() {}