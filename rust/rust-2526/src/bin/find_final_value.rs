struct Solution {}

impl Solution {
    pub fn find_final_value(nums: Vec<i32>, original: i32) -> i32 {
        for num in &nums {
            if *num == original {
                return Self::find_final_value(nums, original * 2);
            }
        }
        original
    }
}

fn main() {
   let ans = Solution::find_final_value(vec![5,3,6,1,12], 3);
    println!("{ans}");
}