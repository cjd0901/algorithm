// Generate Parentheses
// leetcode: https://leetcode-cn.com/problems/generate-parentheses/
struct Solution {}

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut bytes: Vec<char> = vec![];
        let mut ans: Vec<String> = vec![];
        backtrace(0, 0, n, &mut ans, &mut bytes);
        ans
    }
}

pub fn backtrace(left: i32, right: i32, n: i32, ans: &mut Vec<String>, bytes: &mut Vec<char>) {
    if left + right == n * 2 {
        ans.push(bytes.iter().collect());
        return
    }
    if left < n {
        bytes.push('(');
        backtrace(left + 1, right, n, ans, bytes);
        bytes.pop();
    }
    if left > right {
        bytes.push(')');
        backtrace(left, right + 1, n, ans, bytes);
        bytes.pop();
    }
}

fn main() {
   let ans = Solution::generate_parenthesis(3);
    println!("{:?}", ans);
}