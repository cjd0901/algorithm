struct Solution {}

impl Solution {
    pub fn is_zero_array(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> bool {
        let n = nums.len();
        let mut diff = vec![0; n+1];
        for query in queries.iter() {
            diff[query[0] as usize] += 1;
            diff[query[1] as usize + 1] -= 1;
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
}

fn main() {}