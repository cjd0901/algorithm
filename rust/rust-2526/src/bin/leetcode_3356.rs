struct Solution {}

impl Solution {
    fn is_zero_array(nums: &Vec<i32>, queries: &Vec<Vec<i32>>, k: usize) -> bool {
        let n = nums.len();
        let mut diff = vec![0; n+1];
        for i in 0..k {
            let query = &queries[i];
            diff[query[0] as usize] += query[2];
            diff[query[1] as usize+1] -= query[2];
        }

        let mut sum_d = 0;
        for i in 0..n {
            sum_d += diff[i];
            if nums[i] - sum_d > 0 {
                return false;
            }
        }
        true
    }
    pub fn min_zero_array(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> i32 {
        let mut ans = -1;
        let n = queries.len();
        let (mut left, mut right) = (0, n as i32);
        while left <= right {
            let mid = left + (right - left) / 2;
            if Self::is_zero_array(&nums, &queries, mid as usize) {
                ans = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        ans
    }
}

fn main() {
    let nums = vec![4, 3, 2, 1];
    let queries = vec![vec![1, 3, 2], vec![0, 2, 1]];
    let ans = Solution::min_zero_array(nums, queries);
    println!("{}", ans);
}