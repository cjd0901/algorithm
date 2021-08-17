// Student Attendance Record I
// leetcode: https://leetcode-cn.com/problems/student-attendance-record-i/
struct Solution {}
impl Solution {
    pub fn check_record(s: String) -> bool {
        let (mut a, mut l) = (0, 0);
        for ch in s.chars() {
            if ch == 'A' {
                a += 1;
                if a > 1 {
                    return false;
                }
            }
            if ch == 'L' {
                l += 1;
                if l >= 3 {
                    return false;
                }
            } else {
                l = 0;
            }
        }
        true
    }
}

fn main() {
    let ans = Solution::check_record(String::from("PPALLL"));
    println!("{}", ans);
}