// Minimum Moves to Equal Array Elements II
// leetcode: https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/
struct Solution {}
impl Solution {
    pub fn min_moves2(nums: Vec<i32>) -> i32 {
        let mut nums = nums;
        let mut ans = 0;
        nums.sort();
        let mid = nums[nums.len()/2];
        for n in nums.iter() {
            ans += (*n-mid).abs();
        }
        ans
    }
}

fn main() {
    let ans = Solution::min_moves2(vec![1,10,2,9]);
    println!("{}", ans);
}