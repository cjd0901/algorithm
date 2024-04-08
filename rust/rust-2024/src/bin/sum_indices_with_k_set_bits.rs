struct Solution {}

impl Solution {
    pub fn sum_indices_with_k_set_bits(nums: Vec<i32>, k: i32) -> i32 {
        let mut sum = 0;
        for i in 0..nums.len() {
            if i.count_ones() == k as u32 {
                sum += nums[i];
            }
        }
        sum
    }
}

fn main() {}