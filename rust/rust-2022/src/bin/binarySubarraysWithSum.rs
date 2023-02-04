// Binary Subarrays With Sum
// leetcode: https://leetcode-cn.com/problems/binary-subarrays-with-sum/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn num_subarrays_with_sum(nums: Vec<i32>, goal: i32) -> i32 {
        let mut sum = 0;
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        let mut ans = 0;
        for n in nums.iter() {
            {
                *cnt.entry(sum).or_insert(0) += 1;
            }
            sum += *n;
            if let Some(count) = cnt.get(&(sum-goal)) {
                ans += *count;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::num_subarrays_with_sum(vec![1,0,1,0,1], 2);
    println!("{}", ans);
}