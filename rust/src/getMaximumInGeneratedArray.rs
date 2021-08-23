// Get Maximum in Generated Array
// leetcode: https://leetcode-cn.com/problems/get-maximum-in-generated-array/
struct Solution {}
impl Solution {
    pub fn get_maximum_generated(n: i32) -> i32 {
        let mut nums = vec![0; (n+1) as usize];
        let mut ans = 0;
        for i in 0..(n+1) as usize {
            if i == 1 {
                nums[i] = 1;
            }
            if i < ((n+1)/2) as usize {
                nums[i*2] = nums[i];
                nums[i*2+1] = nums[i] + nums[i+1];
            }
            ans = ans.max(nums[i]);
        }
        ans
    }
}

fn main() {
    let ans = Solution::get_maximum_generated(3);
    println!("{}", ans);
}