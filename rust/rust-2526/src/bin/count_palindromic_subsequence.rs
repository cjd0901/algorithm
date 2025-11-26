use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn count_palindromic_subsequence(s: String) -> i32 {
        let mut ans = 0;
        for c in 'a'..='z' {
            let mut chars = s.chars();
            let l = chars.position(|x| x == c);
            let r = chars.rev().position(|x| x == c).map(|x| s.len() - 1 -x);
            if let (Some(l), Some(r)) = (l, r) {
                if r - l < 2 {
                    continue;
                }
                let set = s[l+1..r].chars().collect::<HashSet<char>>();
                ans += set.len() as i32;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::count_palindromic_subsequence("tlpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp".to_string());
    println!("{}", ans);
}