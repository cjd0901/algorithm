// Rabbits in Forest
// leetcode: https://leetcode-cn.com/problems/rabbits-in-forest/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn num_rabbits(answers: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut ans = 0;
        for a in answers.iter() {
            *map.entry(*a).or_insert(0) += 1;
        }
        for (k, v) in map.iter() {
            ans += (k + v) / (k + 1) * (k + 1);
        }
        ans
    }
}

fn main() {
    let ans = Solution::num_rabbits(vec![1,1,2,2,3]);
    println!("{:?}", ans);
}