// Valid Perfect Square
// leetcode: https://leetcode-cn.com/problems/valid-perfect-square/
struct Solution {}

impl Solution {
    pub fn is_perfect_square(num: i32) -> bool {
        let (mut left, mut right) = (0, num);
        while left <= right {
            let mid = (left + (right - left) / 2) as i64;
            let square = mid * mid;
            if square == num as i64 {
                return true;
            } else if square < num as i64 {
                left = mid as i32 + 1;
            } else {
                right = mid as i32 - 1;
            }
        }
        false
    }
}

fn main() {
    let ans = Solution::is_perfect_square(2147483647);
    println!("{}", ans);
}
