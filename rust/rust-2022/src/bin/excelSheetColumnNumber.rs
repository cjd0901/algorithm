// Excel Sheet Column Number
// leetcode： https://leetcode-cn.com/problems/excel-sheet-column-number/
struct Solution {}
impl Solution {
    pub fn title_to_number(column_title: String) -> i32 {
        let titles: Vec<char> = column_title.chars().collect();
        let mut ans = 0;
        for c in titles.iter() {
            ans *= 26;
            ans += (*c as u8 - 'A' as u8) as i32 + 1;
        }
        ans
    }
    pub fn title_to_number2(column_title: String) -> i32 {
        column_title.as_bytes().iter().fold(0, |f, v| {
            f * 26 + (v - b'@') as i32
        })
    }
}

fn main() {
    let ans = Solution::title_to_number2(String::from("ZY"));
    println!("{}", ans);
}