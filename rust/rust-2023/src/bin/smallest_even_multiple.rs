struct Solution {}

// 2413. 最小偶倍数
impl Solution {
    pub fn smallest_even_multiple(n: i32) -> i32 {
        n << (n & 1)
    }
}

fn main() {}
