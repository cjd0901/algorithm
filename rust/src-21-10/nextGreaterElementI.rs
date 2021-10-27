// Next Greater Element I
// leetcode: https://leetcode-cn.com/problems/next-greater-element-i/
use std::collections::HashMap;
struct Solution {}

impl Solution {
    pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut map:HashMap<i32, i32> = HashMap::new();
        let mut stack = vec![];
        let mut ans = vec![-1;nums1.len()];
        for n in nums2.iter().rev() {
            while stack.len() > 0 && n >= stack[stack.len() - 1] {
                stack.pop();
            }
            if stack.len() > 0 {
                map.insert(*n, *stack[stack.len() - 1]);
            } else {
                map.insert(*n, -1);
            }
            stack.push(n);
        }
        for (i, n) in nums1.iter().enumerate() {
            if let Some(a) = map.get(n) {
                ans[i] = *a;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::next_greater_element(vec![4, 1, 2], vec![1, 3, 4, 2]);
    println!("{:?}", ans);
}