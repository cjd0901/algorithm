// Trapping Rain Water
// leetcode: https://leetcode-cn.com/problems/trapping-rain-water/
use std::cmp::min;
struct Solution {}

impl Solution {
    // 单调栈
    pub fn trap(height: Vec<i32>) -> i32 {
        let len = height.len();
        let mut stack = vec![];
        let mut ans = 0;
        for i in 0..len {
            while stack.len() > 0 && height[i] > height[stack[stack.len() - 1]] {
                let h = height[stack.pop().unwrap()];
                if stack.len() == 0 {break;}
                let width = i - stack[stack.len() - 1] - 1;
                let min_height = min(height[stack[stack.len() - 1]], height[i]) - h;
                ans += width as i32 * min_height;
            }
            stack.push(i);
        }
        ans
    }

    // dp
    pub fn trap2(height: Vec<i32>) -> i32 {
        let mut ans = 0;
        let len = height.len();
        let (mut left_max, mut right_max) = (vec![0;len], vec![0;len]);
        let mut cur_max = 0;
        for i in 0..len {
            cur_max = cur_max.max(height[i]);
            left_max[i] = cur_max;
        }
        cur_max = 0;
        for i in (0..len).rev() {
            cur_max = cur_max.max(height[i]);
            right_max[i] = cur_max;
        }
        for i in 0..len {
            let min_height = std::cmp::min(left_max[i], right_max[i]);
            let cur_height = height[i];
            if min_height > cur_height {
                ans += min_height - cur_height;
            }
        }
        ans
    }

    // 双指针
    pub fn trap3(height: Vec<i32>) -> i32 {
        let mut ans = 0;
        let len = height.len();
        let (mut left_max, mut right_max) = (0, height[len-1]);
        let (mut left, mut right) = (0 as usize, len-1);
        while left < right {
            left_max = left_max.max(height[left]);
            right_max = right_max.max(height[right]);
            if left_max < right_max {
                ans += left_max - height[left];
                left += 1;
            } else {
                ans += right_max - height[right];
                right -= 1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::trap3(vec![0,1,0,2,1,0,1,3,2,1,2,1]);
    println!("{}", ans);
}