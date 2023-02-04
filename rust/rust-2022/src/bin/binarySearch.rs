// Binary Search
// leetcode: https://leetcode-cn.com/problems/binary-search/
struct Solution {}
impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let (mut l, mut r) = (0, (nums.len() - 1) as i32);
        while l <= r {
            let mid = l + (r - l) / 2;
            let n = nums[mid as usize];
            if n == target {
                return mid;
            } else if n < target {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
        -1
    }
}

fn main() {
    let ans = Solution::search(vec![1], -1);
    println!("{}", ans);
}