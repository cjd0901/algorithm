// Valid Parentheses
// leetcode: https://leetcode-cn.com/problems/valid-parentheses/
struct Solution {}

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for ch in s.chars() {
            match ch {
                '(' | '[' | '{' => stack.push(ch),
                ']' => {if stack.pop().unwrap_or('0') != '[' {return false}},
                ')' => {if stack.pop().unwrap_or('0') != '(' {return false}},
                '}' => {if stack.pop().unwrap_or('0') != '{' {return false}},
                _ => (),
            }
        }
        stack.len() == 0
    }
}

fn main() {
    let ans = Solution::is_valid(String::from(")"));
    println!("{}", ans);
}