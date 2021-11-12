// Contains Duplicate II
// leetcode: https://leetcode-cn.com/problems/contains-duplicate-ii/
use std::collections::HashSet;
struct Solution {}

impl Solution {
    pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
        let len = nums.len();
        for i in 0..len {
            let end = len.min(i+k as usize+1);
            for j in i..end {
                if j > i && nums[i] == nums[j] {
                    return true;
                }
            }
        }
        false
    }

    pub fn contains_nearby_duplicate2(nums: Vec<i32>, k: i32) -> bool {
        let mut set: HashSet<i32> = HashSet::new();
        let len = nums.len();
        for i in 0..len {
            if set.contains(&nums[i]) { return true; }
            set.insert(nums[i]);
            if set.len() as i32 > k {
                set.remove(&nums[i - k as usize]);
            }
        }
        false
    }
}

fn main() {
    let ans = Solution::contains_nearby_duplicate2(vec![1,2,3,1,2,3], 2);
    println!("{}", ans);
}