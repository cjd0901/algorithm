// Sort Colors
// leetcode: https://leetcode-cn.com/problems/sort-colors/
struct Solution {}

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let (mut red, mut white) = (0 as usize, 0 as usize);
        for i in 0..nums.len() {
            if nums[i] == 0 {
                let temp = nums[red];
                nums[red] = nums[i];
                nums[i] = temp;
                red += 1;
                if nums[i] == 1 {
                    let temp = nums[white];
                    nums[white] = nums[i];
                    nums[i] = temp;
                }
                white += 1;
            } else if nums[i] == 1 {
                let temp = nums[white];
                nums[white] = nums[i];
                nums[i] = temp;
                white += 1;
            }
        }
    }
}

fn main() {
    let mut nums = vec![2];
    Solution::sort_colors(&mut nums);
    println!("{:?}", nums);
}