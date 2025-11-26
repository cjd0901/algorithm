struct Solution {}

impl Solution {
    pub fn minimum_operations(nums: Vec<i32>) -> i32 {
        nums.iter().filter(|&x| x % 3 != 0).count() as i32
    }
}

fn main () {
    let ans = Solution::minimum_operations(vec![1, 2, 3, 4]);
    println!("{ans}");
}