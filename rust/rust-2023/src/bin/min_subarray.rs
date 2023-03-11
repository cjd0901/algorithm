struct Solution {}

// 1590. 使数组和能被 P 整除
impl Solution {
    pub fn min_subarray(nums: Vec<i32>, p: i32) -> i32 {
        let len = nums.len();
        let mut nums = nums;
        nums.insert(0, 0);
        for i in 1..=len {
            nums[i] = (nums[i - 1] + nums[i]) % p;
        }
        let x = nums[len];
        if x == 0 {
            return 0;
        };
        let mut ans = len;
        let mut last = std::collections::HashMap::new();
        for (i, v) in nums.iter().enumerate() {
            *last.entry(v).or_insert(0) = i;
            let key = (*v - x + p) % p;
            if last.contains_key(&key) {
                ans = std::cmp::min(ans, i - last.get(&key).unwrap());
            }
        }
        if ans < len {
            return ans as i32;
        }
        -1
    }
}

fn main() {}
