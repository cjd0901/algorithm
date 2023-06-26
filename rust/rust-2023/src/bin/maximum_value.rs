struct Solution {}

impl Solution {
    pub fn maximum_value(strs: Vec<String>) -> i32 {
        strs.iter().fold(0, |ans, s| {
            ans.max(s.parse::<i32>().unwrap_or(s.len() as i32))
        })
    }
}

fn main() {}
