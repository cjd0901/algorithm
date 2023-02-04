// Search in Rotated Sorted Array
// leetcode: https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let n = nums.len();
        let (mut l, mut r) = (0 as usize, n - 1);
        while l <= r {
            let mid = l + (r - l) / 2;
            if nums[mid] == target {
                return mid as i32;
            }
            if nums[0] <= nums[mid] {
                if target >= nums[0] && target < nums[mid] {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            } else {
                if target > nums[mid] && target <= nums[n-1] {
                    l = mid + 1;
                } else {
                    r = mid - 1;
                }
            }
        }
        -1
    }
}

fn main() {
    let ans = Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 3);
    println!("{}", ans);
}