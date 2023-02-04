// Distribute Candies
// leetcode: https://leetcode-cn.com/problems/distribute-candies/
use std::collections::HashMap;
struct Solution {}

impl Solution {
    pub fn distribute_candies(candy_type: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut left = candy_type.len() / 2;
        let mut map: HashMap<i32, bool> = HashMap::new();
        for i in 0..candy_type.len() {
            if !map.contains_key(&candy_type[i]) && left > 0 {
                map.insert(candy_type[i], true);
                left -= 1;
                ans += 1;
            }
        }
        println!("{}", map.len());
        ans
    }
}

fn main() {
    let ans = Solution::distribute_candies(vec![1,1,2,3]);
    println!("{}", ans);
}