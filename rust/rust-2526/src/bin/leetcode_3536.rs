struct Solution {}

impl Solution {
    pub fn max_product(mut n: i32) -> i32 {
        let mut nums: Vec<i32> = vec![];
        while n > 0 {
            nums.push(n % 10);
            n /= 10;
        }
        nums.sort_by(|a, b| b.cmp(a));
        nums[0] * nums[1]
    }
}

fn main() {}