// Arranging Coins
// leetcode: https://leetcode-cn.com/problems/arranging-coins/
struct Solution {}

impl Solution {
    pub fn arrange_coins(n: i32) -> i32 {
        let (mut lo, mut hi, n) = (0 as u64, n as u64, n as u64);
        let mut max: u64 = 0;
        while lo <= hi {
            let mid = lo + (hi - lo) / 2;
            if mid * mid + mid > 2 * n {
                hi = mid - 1;
            } else {
                max = mid;
                lo = mid + 1;
            }
        }
        max as i32
    }
}

fn main() {
    let max = Solution::arrange_coins(1804289383);
    println!("{}", max);
}