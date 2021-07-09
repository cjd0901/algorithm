// Find Majority Element
// leetcode: https://leetcode-cn.com/problems/find-majority-element-lcci/
struct Solution {}
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut candidate = -1;
        let mut count = 0;
        for n in nums.iter() {
            if count == 0 {
                candidate = *n;
            }
            if *n == candidate {
                count += 1;
            } else {
                count -= 1;
            }
        }
        count = 0;
        for n in nums.iter() {
            if *n == candidate {
                count += 1;
            }
        }
        if count > nums.len() / 2 {
            return candidate;
        }
        -1
    }
}

fn main() {
    let ans = Solution::majority_element(vec![1,2,5,9,5,9,5,5,5]);
    println!("{}", ans);
}