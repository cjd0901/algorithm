use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn repeated_n_times(nums: Vec<i32>) -> i32 {
        let mut ans = -1;
        let n = nums.len() / 2;
        let mut m = HashMap::new();
        for num in nums {
            let c = m.entry(num).or_insert(0);
            *c += 1;
            if *c == n {
                ans = num;
                break;
            }
        }
        ans
    }
}

fn main() {

}