// Count Good Meals
// https://leetcode-cn.com/problems/count-good-meals/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn count_pairs(deliciousness: Vec<i32>) -> i32 {
        let mut max_val = 0;
        let mut ans = 0;
        for v in deliciousness.iter() {
            if *v > max_val {
                max_val = *v;
            }
        }
        let max_sum = max_val * 2;
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        for v in deliciousness.iter() {
            let mut sum = 1;
            while sum <= max_sum {
                if let Some(count) = cnt.get(&(sum-*v)) {
                    ans = (ans + count) % (1e9 as i32 + 7);
                }
                sum <<= 1;
            }
            let c = cnt.entry(*v).or_insert(0);
            *c += 1;
        }
        ans
    }
}

fn main() {
    let ans = Solution::count_pairs(vec![1,3,5,7,9]);
    println!("{}", ans);
}