// Buddy Strings
// leetcode: https://leetcode-cn.com/problems/buddy-strings/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn buddy_strings(s: String, goal: String) -> bool {
        let mut map: HashMap<u8, i32> = HashMap::new();
        let (s_bytes, goal_bytes) = (s.as_bytes(), goal.as_bytes());
        if s_bytes.len() != goal_bytes.len() {
            return false;
        }
        if s == goal {
            for b in s_bytes.iter() {
                let count = map.entry(*b).or_insert(0);
                *count += 1;
                if *count > 1 {
                    return true
                }
            }
        } else {
            let mut count = 0;
            let (mut p, mut q) = (-1, -1);
            for i in 0..s_bytes.len() {
                if s_bytes[i] != goal_bytes[i] {
                    count += 1;
                    if p == -1 {
                        p = i as i32;
                    } else if q == -1 {
                        q = i as i32;
                    }
                }
            }
            if count == 2 {
                return s_bytes[p as usize] == goal_bytes[q as usize] && s_bytes[q as usize] == goal_bytes[p as usize];
            }
        }
        false
    }
}

fn main() {
    let ans = Solution::buddy_strings(String::from("abcaa"), String::from("abcbb"));
    println!("{}", ans);
}