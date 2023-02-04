// Largest Rectangle in Histogram
// leetcode: https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
struct Solution {}

impl Solution {
    pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
        let len = heights.len();
        if len == 1 {
            return heights[0];
        }
        let mut stack: Vec<usize> = Vec::new();
        let mut ans = 0;
        for i in 0..len {
            while stack.len() > 0 && heights[i] < heights[stack[stack.len() - 1]] {
                let height = heights[stack.pop().unwrap()];
                while stack.len() > 0 && height == heights[stack[stack.len() - 1]] {
                    stack.pop();
                }

                let width;
                if stack.len() == 0 {
                    width = i as i32;
                } else {
                    width = (i - stack[stack.len() - 1] - 1) as i32;
                }
                ans = ans.max(width * height);
            }
            stack.push(i);
        }
        while stack.len() > 0 {
            let height = heights[stack.pop().unwrap()];
            while stack.len() > 0 && height == heights[stack[stack.len() - 1]] {
                stack.pop();
            }

            let width;
            if stack.len() == 0 {
                width = len as i32;
            } else {
                width = (len - stack[stack.len() - 1] - 1) as i32;
            }
            ans = ans.max(width * height);
        }
        ans
    }

    pub fn largest_rectangle_area2(heights: Vec<i32>) -> i32 {
        let mut len = heights.len();
        if len == 1 {
            return heights[0];
        }
        let mut heights = heights;
        let mut ans = 0;
        heights.insert(0, 0);
        heights.push(0);
        len = heights.len();

        let mut stack = vec![0 as usize];
        for i in 1..len {
            while heights[i] < heights[stack[stack.len() - 1]] {
                let height = heights[stack.pop().unwrap()];
                let width = (i - stack[stack.len() - 1] - 1) as i32;
                ans = ans.max(height * width);
            }
            stack.push(i);
        }
        ans
    }
}

fn main() {
    let ans = Solution::largest_rectangle_area2(vec![2,1,5,6,2,3]);
    println!("{}", ans);
}