// Max Consecutive Ones
// https://leetcode-cn.com/problems/max-consecutive-ones/
struct Solution {}
impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut count = 0;
        for n in nums.iter() {
            if *n == 1 {
                count += 1;
            } else {
                ans = ans.max(count);
                count = 0;
            }
        }
        ans = ans.max(count);
        ans
    }
}

fn main() {
    let ans = Solution::find_max_consecutive_ones(vec![1,1,0,1,1,1]);
    println!("{}", ans);
}