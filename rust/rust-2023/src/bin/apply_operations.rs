struct Solution {}

impl Solution {
    pub fn apply_operations(mut nums: Vec<i32>) -> Vec<i32> {
        let mut j = 0;
        for i in 0..nums.len() {
            if i < nums.len() - 1 && nums[i] == nums[i + 1] {
                nums[i] *= 2;
                nums[i + 1] = 0;
            }
            if nums[i] != 0 {
                let temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
                j += 1;
            }
        }
        nums
    }
}

fn main() {}
