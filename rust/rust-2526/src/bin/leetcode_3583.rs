use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn special_triplets(nums: Vec<i32>) -> i32 {
        let mut ans: i64 = 0;
        let mut suf: HashMap<i32, i32> = HashMap::new();
        for &num in &nums {
            *suf.entry(num).or_insert(0) += 1;
        }

        let mut pre: HashMap<i32, i32> = HashMap::new();
        for &num in &nums {
            *suf.entry(num).or_insert(0) -= 1;
            ans += *suf.get(&(num*2)).unwrap_or(&0) as i64 *
                *pre.get(&(num*2)).unwrap_or(&0) as i64;
            *pre.entry(num).or_insert(0) += 1;
        }

        (ans % 1_000_000_007) as i32
    }
}

fn main() {

}