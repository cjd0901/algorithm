// Bag of Tokens
// https://leetcode-cn.com/problems/bag-of-tokens/
struct Solution {}
impl Solution {
    pub fn bag_of_tokens_score(tokens: Vec<i32>, power: i32) -> i32 {
        if tokens.len() == 0 {
            return 0;
        }
        let mut tokens = tokens;
        let mut power = power;
        let (mut score, mut ans) = (0, 0);
        let (mut left, mut right) = (0, tokens.len()-1);
        tokens.sort();
        while left <= right && (power >= tokens[left] || score > 0) {
            if power >= tokens[left] {
                score += 1;
                power -= tokens[left];
                left += 1;
            }else {
                score -= 1;
                power += tokens[right];
                right -= 1;
            }
            ans = ans.max(score);
        }
        ans
    }
}

fn main() {
    let score = Solution::bag_of_tokens_score(vec![100, 200], 150);
    println!("{}", score);
}