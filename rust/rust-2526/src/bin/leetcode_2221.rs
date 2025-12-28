struct Solution {}

impl Solution {
    pub fn triangular_sum(mut nums: Vec<i32>) -> i32 {
        let mut n = nums.len();
        while n > 1 {
            for i in 0..n-1 {
                nums[i] = (nums[i] + nums[i+1]) % 10;
            }
            n -= 1;
        }
        nums[0]
    }
}

fn main() {}