use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn find_final_value(nums: Vec<i32>, k: i32, x: i32) -> Vec<i32> {
        let len = nums.len();
        let mut ans: Vec<i32> = vec![0; len-k as usize+1];
        let mut i = 0;
        while (i+k as usize) <= len {
            let mut map: HashMap<i32, i32> = HashMap::new();
            for num in nums[i..i+k as usize].iter() {
                *map.entry(*num).or_insert(0) += 1;
            }
            let mut v = map.into_iter().collect::<Vec<(i32, i32)>>();
            v.sort_by(|a, b| b.1.cmp(&a.1).then_with(|| b.0.cmp(&a.0)));
            let l = v.len().min(x as usize);
            for m in 0..l as usize {
                ans[i] += v[m].0 * v[m].1;
            }
            i += 1;
        }
        ans
    }
}

fn main() {
    Solution::find_final_value(vec![9, 2, 2], 3, 3);
}