struct Solution {}

impl Solution {
    pub fn prefixes_div_by5(nums: Vec<i32>) -> Vec<bool> {
        let mut ans = vec![false; nums.len()];
        let mut x = 0;
        for (i, num) in nums.iter().enumerate() {
            x = (x << 1 | num) % 5;
            ans[i] = x == 0;
        }
        ans
    }
}

fn main() {
    Solution::prefixes_div_by5(vec![1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0]);
}