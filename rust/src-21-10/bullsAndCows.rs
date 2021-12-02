// Bulls and Cows
// leetcode: https://leetcode-cn.com/problems/bulls-and-cows/
use std::collections::HashMap;
struct Solution {}

impl Solution {
    pub fn get_hint(secret: String, guess: String) -> String {
        let secret = secret.as_bytes();
        let guess = guess.as_bytes();
        let (mut bulls, mut cows) = (0, 0);
        let mut cow1: HashMap<u8, i32> = HashMap::new();
        let mut cow2: HashMap<u8, i32> = HashMap::new();
        for i in 0..secret.len() {
            if secret[i] == guess[i] {
                bulls += 1;
            } else {
                *cow1.entry(secret[i]).or_insert(0) += 1;
                *cow2.entry(guess[i]).or_insert(0) += 1;
            }
        }
        for (key, value) in cow2.iter() {
            cows += value.min(cow1.get(key).unwrap_or(&0));
        }
        format!("{}A{}B", bulls, cows)
    }
}

fn main() {
    let ans = Solution::get_hint(String::from("1807"), String::from("7801"));
    println!("{}", ans);
}