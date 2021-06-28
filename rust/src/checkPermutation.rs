// Check Permutation
// leetcode: https://leetcode-cn.com/problems/check-permutation-lcci/
use::std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn check_permutation(s1: String, s2: String) -> bool {
        let mut count: HashMap<char, i32> = HashMap::new();
        for s in s1.chars() {
            let count = count.entry(s).or_insert(0);
            *count += 1;
        }
        for s in s2.chars() {
            let count = count.entry(s).or_insert(0);
            *count -= 1;
        }
        for (_k, v) in &count {
            if *v != 0 {
                return false
            }
        }
        true
    }
}

fn main() {
    let ans = Solution::check_permutation(String::from("abc"), String::from("bca"));
    println!("{}", ans);
}