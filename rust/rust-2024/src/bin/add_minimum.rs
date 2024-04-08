struct Solution {}

impl Solution {
    pub fn add_minimum(word: String) -> i32 {
        let len = word.len();
        let word: Vec<char> = word.chars().collect::<_>();
        let mut dp = vec![0; len + 1];
        for i in 1..=len {
            dp[i] = dp[i - 1] + 2;
            if i > 1 && word[i - 1] > word[i - 2] {
                dp[i] = dp[i - 1] - 1;
            }
        }
        dp[len]
    }
}

fn main () {}