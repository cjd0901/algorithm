// Number of Substrings with only 1s
// https://leetcode-cn.com/problems/number-of-substrings-with-only-1s/
struct Solution {}
impl Solution {
    pub fn num_sub(s: String) -> i32 {
        let mut count:u64 = 0;
        let mut total = 0;
        const M: u64 = 1e9 as u64 + 7;
        for c in s.chars() {
            if c == '1' {
                count += 1;
            }else if c == '0' {
                total += count * (count + 1) / 2;
                total = total % M;
                count = 0;
            }
        }
        total += count * (count + 1) / 2;
        total = total % M;
        total as i32
    }
}

fn main() {
    let ans = Solution::num_sub(String::from("111111111111111111111111111"));
    println!("{}", ans);
}