// Plus One
// leetcode: https://leetcode-cn.com/problems/plus-one/
struct Solution {}

impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = Vec::new();
        let mut carry = 1;
        for digit in digits.iter().rev() {
            if carry == 1 {
                if *digit == 9 {
                    ans.push(0);
                } else {
                    ans.push(*digit+carry);
                    carry = 0;
                }
            } else {
                ans.push(*digit);
            }
        }
        if carry > 0 {
            ans.push(carry);
        }
        ans.reverse();
        ans
    }
}

fn main() {
    let ans = Solution::plus_one(vec![9, 9, 9]);
    println!("{:?}", ans);
}