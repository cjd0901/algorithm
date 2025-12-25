struct Solution {}

impl Solution {
    pub fn count_valid_selections(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = nums.len();
        let non_zero_count = nums.iter().filter(|&&x| x != 0).count();
        for i in 0..n {
            if nums[i] == 0 {
                if Self::is_valid(&nums, non_zero_count, i as i32, 1) { ans += 1; }
                if Self::is_valid(&nums, non_zero_count, i as i32, -1) { ans += 1; }
            }
        }
        ans
    }

    fn is_valid(nums: &Vec<i32>, mut non_zero_count: usize, mut ptr: i32, mut direction: i32) -> bool {
        let mut temp = nums.clone();
        while ptr >= 0 && ptr <= nums.len() as i32 - 1 {
            if temp[ptr as usize] > 0 {
                temp[ptr as usize] -= 1;
                direction *= -1;
                if temp[ptr as usize] == 0 {
                    non_zero_count -= 1;
                }
            }
            ptr += direction;
        }
        non_zero_count == 0
    }
}

fn main() {}