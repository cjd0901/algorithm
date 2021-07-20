// Minimize Maximum Pair Sum in Array
// leetcode: https://leetcode-cn.com/problems/minimize-maximum-pair-sum-in-array/
struct Solution {}
impl Solution {
    pub fn min_pair_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len()-1;
        let mut nums = nums;
        nums.sort();
        let mut ans = nums[0];
        for i in 0..nums.len()/2 {
            ans = ans.max(nums[i] + nums[n-i]);
        }
        ans
    }
}

fn main() {
    let ans = Solution::min_pair_sum(vec![3,5,2,3]);
    println!("{}", ans);
}