// Number of Subarrays with Bounded Maximum
// leetCode: https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum/
struct Solution {}
impl Solution {
    pub fn num_subarray_bounded_max(nums: Vec<i32>, left: i32, right: i32) -> i32 {
        let n = nums.len();
        let (mut last_break, mut last_count) = (-1, 0);
        let mut total = 0;
        for i in 0..n {
            let num = nums[i];
            if num > right {
                last_break = i as i32;
                last_count = 0;
            }else if num >= left {
                last_count = i as i32 - last_break;
                total += last_count;
            }else {
                total += last_count;
            }
        }
        total
    }
}

fn main() {
    let num = Solution::num_subarray_bounded_max(vec![73,55,36,5,55,14,9,7,72,52], 32, 69);
    println!("{}", num);
}