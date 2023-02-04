// Split a String in Balanced Strings
// leetcode: https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
struct Solution {}
impl Solution {
    pub fn balanced_string_split(s: String) -> i32 {
        let (mut l, mut r) = (0, 0);
        let mut ans = 0;
        for ch in s.chars() {
            if ch == 'R' {
                if l > 0 {
                    l -= 1;
                    if l == 0 {
                        ans += 1;
                    }
                } else {
                    r += 1;
                }
            } else {
                if r > 0 {
                    r -= 1;
                    if r == 0 {
                        ans += 1;
                    }
                } else {
                    l += 1;
                }
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::balanced_string_split(String::from("RLRRLLRLRL"));
    println!("{}", ans);
}
