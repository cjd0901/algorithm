// Frequency of the Most Frequent Element
// leetcode: https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element/
struct Solution {}
impl Solution {
    pub fn max_frequency(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let (mut l, mut ans) = (0, 1);
        let n = nums.len();
        let mut total = 0;
        nums.sort();
        for r in 1..n {
            total += (nums[r] - nums[r-1]) * (r as i32 - l);
            while total > k {
                total -= nums[r] - nums[l as usize];
                l += 1;
            }
            ans = ans.max(r as i32-l+1);
        }
        ans as i32
    }
}

fn main() {
    let ans = Solution::max_frequency(vec![1,4,8,13], 5);
    println!("{}", ans);
}