// Single Number
// leetcode: https://leetcode-cn.com/problems/single-number-iii/
struct Solution {}
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        let mut xor_sum = 0;
        for num in nums.iter() {
            xor_sum ^= num;
        }
        let mut temp = 1;
        while xor_sum & temp == 0 {
            temp <<= 1;
        }
        let (mut a, mut b) = (0, 0);
        for num in nums.iter() {
            if num & temp != 0 {
                a ^= num;
            }else {
                b ^= num;
            }
        }
        vec![a, b]
    }
}

fn main() {
    let ans = Solution::single_number(vec![-1,0]);
    println!("{:?}", ans);
    println!("{}", 6 & -6);
}