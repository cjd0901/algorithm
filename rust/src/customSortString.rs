// Custom Sort String
// leetcode: https://leetcode-cn.com/problems/custom-sort-string/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn custom_sort_string(order: String, s: String) -> String {
        let mut map: HashMap<char, i32> = HashMap::new();
        let mut new_s = String::from("");
        for str in s.chars() {
            *map.entry(str).or_insert(0) += 1;
        }
        for o in order.chars() {
            if let Some(count) = map.get_mut(&o) {
                for _ in 0..*count {
                    new_s.push(o);
                    *count -= 1;
                }
            }
        }
        for (k, v) in map.iter() {
            for _ in 0..*v {
                new_s.push(*k);
            }
        }
        new_s
    }
}

fn main() {
    let ans = Solution::custom_sort_string(String::from("cba"), String::from("abcd"));
    println!("{}", ans);
}