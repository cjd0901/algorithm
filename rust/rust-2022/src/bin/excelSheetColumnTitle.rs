// Excel Sheet Column Title
// leetcode: https://leetcode-cn.com/problems/excel-sheet-column-title/
struct Solution {}
impl Solution {
    pub fn convert_to_title(column_number: i32) -> String {
        let mut ans = String::from("");
        let mut column_number = column_number;
        while column_number > 0 {
            let ch = (column_number - 1) % 26;
            column_number = (column_number-1) / 26;
            ans.push(('A' as u8 + ch as u8) as char);
        }
        ans.chars().rev().collect()
    }
}

fn main() {
    let ans = Solution::convert_to_title(701);
    println!("{}", ans);
}