struct Solution {}

impl Solution {
    // 超时
    pub fn max_sub_array(nums: Vec<i32>, k: i32) -> i64 {
        let mut ans = i64::MIN;
        let mut v: Vec<i64> = vec![0; nums.len() + 1];
        for i in 1..=nums.len() {
            v[i] = v[i - 1] + nums[i-1] as i64;
        }
        for l in 0..=(nums.len() - k as usize) {
            let mut r = l;
            while (r + k as usize) < v.len() {
                r += k as usize;
                ans = ans.max(v[r] - v[l]);
            }
        }
        ans
    }

    pub fn max_sub_array2(nums: Vec<i32>, k: i32) -> i64 {
        let mut ans = i64::MIN;
        let mut v: Vec<i64> = vec![i64::MAX / 2; k as usize];
        let mut prefix_sum: i64 = 0;
        v[k as usize - 1] = 0;
        for i in 0..nums.len() {
            prefix_sum += nums[i] as i64;
            let sum = prefix_sum - v[i % k as usize];
            ans = ans.max(sum);
            v[i % k as usize] = v[i % k as usize].min(prefix_sum);
        }
        ans
    }
}

fn main() {
    let ans = Solution::max_sub_array2(vec![-7, -9], 1);
    println!("{:?}", ans);
}