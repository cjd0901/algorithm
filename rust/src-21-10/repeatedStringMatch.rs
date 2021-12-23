// Repeated String Match
// leetcode: https://leetcode-cn.com/problems/repeated-string-match/
struct Solution {}

impl Solution {
    pub fn repeated_string_match(a: String, b: String) -> i32 {
        let m = a.len();
        let n = b.len();
        for time in (n / m) ..= (n / m + 2) {
            if a.repeat(time).contains(&b) {
                return time as i32;
            }
        }
        -1
    }
}

fn main() {
    let ans = Solution::repeated_string_match(String::from("abcd"), String::from("cdabcdab2"));
    println!("{}", ans);
}