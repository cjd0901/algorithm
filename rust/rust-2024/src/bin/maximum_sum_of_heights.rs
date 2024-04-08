struct Solution {}

impl Solution {
    pub fn maximum_sum_of_heights(max_heights: Vec<i32>) -> i64 {
        let mut sum = 0;
        let len = max_heights.len();
        for i in 0..len {
            let mut temp_sum = max_heights[i] as i64;
            let mut pre_height = max_heights[i];
            for j in (1..=i).rev() {
                if max_heights[j - 1] < pre_height {
                    temp_sum += max_heights[j - 1] as i64;
                    pre_height = max_heights[j - 1];
                } else {
                    temp_sum += pre_height as i64;
                }
            }

            let mut pre_height = max_heights[i];
            for k in i..len - 1 {
                if max_heights[k + 1] < pre_height {
                    temp_sum += max_heights[k + 1] as i64;
                    pre_height = max_heights[k + 1];
                } else {
                    temp_sum += pre_height as i64;
                }
            }
            sum = sum.max(temp_sum);
        }
        sum
    }
}

fn main() {}