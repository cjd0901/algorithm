struct Solution {}

// 2395. 和相等的子数组
impl Solution {
    pub fn find_subarrays(nums: Vec<i32>) -> bool {
        let mut m = std::collections::HashSet::new();
        for i in 0..nums.len() - 1 {
            let sum = nums[i] + nums[i + 1];
            if !m.insert(sum) {
                return true;
            }
        }
        false
    }
}

fn main() {}
