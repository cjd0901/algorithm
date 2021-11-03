// Add Binary
// leetcode: https://leetcode-cn.com/problems/add-binary/
struct Solution {}

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        let (a, b): (Vec<char>, Vec<char>) = (a.chars().collect(), b.chars().collect());
        let mut ans = vec![];
        let max_len = std::cmp::max(a.len(), b.len());
        let mut carry = 0;
        for i in 0..max_len {
            if i < a.len() {
                carry += (a[(a.len() - i - 1)] as u8 - '0' as u8) as i32;
            }
            if i < b.len() {
                carry += (b[(b.len() - i - 1)] as u8 - '0' as u8) as i32;
            }
            ans.push(((carry % 2) as u8 + '0' as u8) as char);
            carry /= 2;
        }
        if carry == 1 {
            ans.push('1');
        }
        ans.reverse();
        ans.iter().collect()
    }
}

fn main() {
    let ans = Solution::add_binary(String::from("1010"), String::from("1011"));
    println!("{}", ans);
}