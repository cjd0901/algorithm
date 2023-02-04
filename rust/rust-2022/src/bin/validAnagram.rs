// Valid Anagram
// leetcode: https://leetcode-cn.com/problems/valid-anagram/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut map: HashMap<char, i32> = HashMap::new();
        if s.len() != t.len() {
            return false
        }
        for c in s.chars() {
            *map.entry(c).or_insert(0) += 1;
        }
        for c in t.chars() {
            *map.entry(c).or_insert(0) -= 1;
        }
        for v in map.values() {
            if *v > 0 {
                return false
            }
        }
        true
    }
}

fn main() {
    let ans = Solution::is_anagram(String::from("rat"), String::from("car"));
    println!("{}", ans);
}