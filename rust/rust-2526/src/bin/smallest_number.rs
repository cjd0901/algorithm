struct Solution {}

impl Solution {
    pub fn smallest_number(n: i32) -> i32 {
        let leading_zeros = n.leading_zeros();
        let effective_bits = i32::BITS - leading_zeros;
        (1 << effective_bits) - 1
    }
}

fn main() {
    let ans = Solution::smallest_number(7);
    println!("{ans}");
}