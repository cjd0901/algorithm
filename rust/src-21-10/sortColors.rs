// Sort Colors
// leetcode: https://leetcode-cn.com/problems/sort-colors/
struct Solution {}

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let (mut red, mut white) = (0 as usize, 0 as usize);
        for i in 0..nums.len() {
            if nums[i] == 0 {
                swap(nums, i, red);
                red += 1;
                if nums[i] == 1 {
                    swap(nums, i, white);
                }
                white += 1;
            } else if nums[i] == 1 {
                swap(nums, i, white);
                white += 1;
            }
        }
    }

    pub fn sort_colors2(nums: &mut Vec<i32>) {
        let len = nums.len();
        // [0, p0) == 0
        // [p0, i) == 1
        // (p2, len-1] == 2
        let (mut p0, mut i, mut p2) = (0 as usize, 0 as usize, (len-1) as i32);
        while i as i32 <= p2 {
            if nums[i] == 0 {
                swap(nums, i, p0);
                p0 += 1;
                i += 1;
            } else if nums[i] == 1 {
                i += 1;
            } else {
                swap(nums, i, p2 as usize);
                p2 -= 1;
            }
        }
    }
}

fn swap(nums: &mut Vec<i32>, i: usize, j: usize) {
    let temp = nums[i];
    nums[i] = nums[j];
    nums[j] = temp;
}

fn main() {
    let mut nums = vec![2];
    Solution::sort_colors2(&mut nums);
    println!("{:?}", nums);
}