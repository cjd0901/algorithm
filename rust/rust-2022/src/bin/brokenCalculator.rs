// Broken Calculator
// leetcode: https://leetcode-cn.com/problems/broken-calculator/
struct Solution {}
impl Solution {
    pub fn broken_calc(x: i32, y: i32) -> i32 {
        let mut count = 0;
        let mut y = y;
        while y > x {
            count += 1;
            if y%2 == 1 {
                y += 1;
            } else {
                y /= 2;
            }
        }
        count + x - y
    }
}

fn main() {
    let ans = Solution::broken_calc(2, 3);
    println!("{}", ans);
}