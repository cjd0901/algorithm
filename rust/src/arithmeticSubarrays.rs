// Arithmetic Subarrays
// leetcode: https://leetcode-cn.com/problems/arithmetic-subarrays/
struct Solution{}

impl Solution {
    pub fn check_arithmetic_subarrays(nums: Vec<i32>, l: Vec<i32>, r: Vec<i32>) -> Vec<bool> {
        let n = l.len();
        let mut ans: Vec<bool> = vec![true; n];
        for i in 0..n {
            let mut temp = nums[l[i] as usize..(r[i]+1) as usize].to_vec();
            temp.sort();
            let sub = temp[1] - temp[0];
            for k in 0..temp.len() - 1 {
                if temp[k+1] - temp[k] != sub {
                    ans[i] = false;
                }
            }
        }
        ans
    }
}

fn main() {
    let nums = vec![4, 6, 5, 9, 3, 7];
    let l = vec![0, 0, 2];
    let r = vec![2, 3, 5];
    let ans = Solution::check_arithmetic_subarrays(nums, l, r);
    println!("{:?}", ans);
}