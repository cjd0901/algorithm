use std::collections::VecDeque;

struct Solution {}

impl Solution {
    pub fn min_length(s: String) -> i32 {
        let mut stack = VecDeque::new();
        for ch in s.chars() {
            match stack.back() {
                Some(&'A') if ch == 'B' => {
                    stack.pop_back();
                },
                Some(&'C') if ch == 'D' => {
                    stack.pop_back();
                },
                _ => {
                    stack.push_back(ch);
                }
            }
        }
        stack.len() as i32
    }
}

fn main() {}