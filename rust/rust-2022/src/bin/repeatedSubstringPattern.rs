// Repeated Substring Pattern
// leetcode: https://leetcode-cn.com/problems/repeated-substring-pattern/
struct Solution {}
impl Solution {
    pub fn repeated_substring_pattern(s: String) -> bool {
        let s = s.as_bytes();
        let n = s.len();
        for i in 1..n/2+1 {
            if n%i == 0 {
                let mut b = true;
                for j in i..n {
                    if s[j] != s[j-i] {
                        b = false;
                        break;
                    }
                }
                if b {
                    return true;
                }
            }
        }
        false
    }
}

fn main() {
    let ans =Solution::repeated_substring_pattern(String::from("abab"));
    println!("{}", ans);
}