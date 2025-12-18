struct Solution {}

impl Solution {
    pub fn arrange_coins(n: i32) -> i32 {
        let (mut left, mut right): (i64, i64) = (1, n as i64);
        while left < right {
            let mid = left + (right - left + 1) / 2;
            if mid * (mid + 1) <= 2 * (n as i64) {
                left = mid;
            } else {
                right = mid - 1;
            }
        }
        left as i32
    }
}

fn main() {}