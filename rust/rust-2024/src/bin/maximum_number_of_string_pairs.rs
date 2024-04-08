use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
        let mut ans = 0;
        let mut map = HashMap::new();
        for word in words.iter() {
            let word_rev = word.chars().rev().collect::<String>();
            if map.contains_key(&word_rev) {
                ans += 1;
            } else {
                map.insert(word, 0);
            }
        }
        ans
    }
}

fn main() {}