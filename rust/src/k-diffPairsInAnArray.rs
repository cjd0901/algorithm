// K-diff Pairs in an Array
// leetcode: https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_pairs(nums: Vec<i32>, k: i32) -> i32 {
        let mut map: HashMap<i32, Vec<i32>> = HashMap::new();
        let mut ans = 0;
        for n in nums.iter() {
            map.entry(*n).or_insert(vec![0; 2])[0] += 1;
        }
        for key in map.clone().keys() {
            if let Some(v) = map.get_mut(&(key+k)) {
                if v[1] == 1 {
                    continue;
                }
                if k == 0 && v[0] > 1 || k > 0 {
                    ans += 1;
                }
                v[1] = 1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::find_pairs(vec![1,1,1,1,1], 0);
    println!("{}", ans);
}