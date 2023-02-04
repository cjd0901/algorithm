// Check if Binary String Has at Most One Segment of Ones
// leetcode: https://leetcode-cn.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
struct Solution {}
impl Solution {
    pub fn check_ones_segment(s: String) -> bool {
        let mut count = 0;
        let mut one = 0;
        for ch in s.chars() {
            if ch == '1' {
                one += 1;
            } else {
                if one > 0 {
                    count += 1;
                }
                one = 0;
            }
        }
        if one > 0 {
            count += 1;
        }
        count == 1
    }
}

fn main() {
    let ans = Solution::check_ones_segment(String::from("1001"));
    println!("{}", ans);
}