use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn count_words(words1: Vec<String>, words2: Vec<String>) -> i32 {
        let mut ans = 0;
        let (mut w1_map, mut w2_map) = (HashMap::new(), HashMap::new());
        for word in words1.into_iter() {
            *w1_map.entry(word).or_insert(0) += 1;
        }

        for word in words2.into_iter() {
            *w2_map.entry(word).or_insert(0) += 1;
        }

        for (k, v) in w1_map.into_iter() {
            if v == 1 && *w2_map.get(&k).unwrap_or(&0) == 1 {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {}