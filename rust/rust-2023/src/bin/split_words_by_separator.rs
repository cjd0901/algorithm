struct Solution {}

impl Solution {
    pub fn split_words_by_separator(words: Vec<String>, separator: char) -> Vec<String> {
        words
            .iter()
            .flat_map(|w| w.split(separator).collect::<Vec<_>>())
            .map(|w| w.to_string())
            .filter(|w| !w.is_empty())
            .collect()
    }
}

fn main() {}