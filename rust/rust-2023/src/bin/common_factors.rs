struct Solution {}

// 2427. 公因子的数目
impl Solution {
    pub fn common_factors(a: i32, b: i32) -> i32 {
        (1..=a.min(b)).fold(0, |cnt, x| {
            if a % x == 0 && b % x == 0 {
                cnt + 1
            } else {
                cnt
            }
        })
    }
}

fn main() {}
