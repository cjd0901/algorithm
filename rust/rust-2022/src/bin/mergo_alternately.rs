struct Solution {}

impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {
        let (mut i, mut j) = (0, 0);
        let mut ans = String::from("");
        while i < word1.len() || j < word2.len() {
            if i < word1.len() {
                ans.push(word1.chars().nth(i as usize).unwrap());
                i += 1;
            }
            if j < word2.len() {
                ans.push(word2.chars().nth(j as usize).unwrap());
                j += 1;
            }
        }
        ans
    }
}

fn main() {
    let word1 = String::from("abc");
    let word2 = String::from("efgh");
    let ans = Solution::merge_alternately(word1, word2);
    println!("{}", ans);
}