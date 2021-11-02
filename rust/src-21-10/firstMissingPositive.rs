// First Missing Positive
// leetcode: https://leetcode-cn.com/problems/first-missing-positive/
struct Solution {}

impl Solution {
    pub fn first_missing_positive(nums: Vec<i32>) -> i32 {
        let len = nums.len();
        let mut nums = nums;
        for i in 0..len {
            while nums[i] > 0 && nums[i] <= len as i32 && nums[i] != nums[(nums[i]-1) as usize] {
                let num = nums[i];
                nums.swap(i, (num - 1) as usize);
            }
        }
        for i in 0..len {
            if nums[i] - 1 != i as i32 {
                return (i + 1) as i32;
            }
        }
        (len + 1) as i32
    }
}

fn main() {
    let ans = Solution::first_missing_positive(vec![1, 2, 0]);
    println!("First missing positive-{:?}", ans)
}