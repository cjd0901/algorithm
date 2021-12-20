// Intersection of Two Arrays II
// leetcode: https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
struct Solution {}

impl Solution {
    pub fn intersect(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let (mut nums1, mut nums2) = (nums1, nums2);
        let mut ans = Vec::new();
        nums1.sort();
        nums2.sort();
        let (mut i, mut j) = (0 as usize, 0 as usize);
        while i < nums1.len() && j < nums2.len() {
            if nums1[i] < nums2[j] {
                i += 1;
            } else if nums1[i] > nums2[j] {
                j += 1;
            } else {
                ans.push(nums1[i]);
                i += 1;
                j += 1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::intersect(vec![1,2,2,1], vec![2,2]);
    println!("{:?}", ans);
}