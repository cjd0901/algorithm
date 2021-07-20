// Number of Good Pairs
// leetcode: https://leetcode-cn.com/problems/number-of-good-pairs/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut num = 0;
        for n in nums.iter() {
            *map.entry(*n).or_insert(0) += 1;
        }
        for v in map.values() {
            if *v > 1 {
                num += *v * (*v-1) / 2;
            }
        }
        num
    }
}

fn main() {
    let ans = Solution::num_identical_pairs(vec![1,2,3,1,1,3]);
    println!("{}", ans);
}