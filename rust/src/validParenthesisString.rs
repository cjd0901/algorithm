// Valid Parenthesis String
// leetcode: https://leetcode-cn.com/problems/valid-parenthesis-string/
struct Solution {}
impl Solution {
    pub fn check_valid_string(s: String) -> bool {
        let chs: Vec<char> = s.chars().collect();
        let l = chs.len();
        let mut dp = vec![vec![false; l]; l];
        for i in 0..l {
            if chs[i] == '*' {
                dp[i][i] = true;
            }
        }
        for i in 1..l {
            let (c1, c2) = (chs[i-1], chs[i]);
            if (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*') {
                dp[i-1][i] = true;
            }
        }
        if l > 2 {
            for i in (0..l-2).rev() {
                let c1 = chs[i];
                for j in i+2..l {
                    let c2 = chs[j];
                    if (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*') {
                        dp[i][j] = dp[i+1][j-1];
                    }
                    let mut k = i;
                    while k <= j && !dp[i][j] {
                        dp[i][j] = dp[i][k] && dp[k+1][j];
                        k += 1;
                    }
                }
            }
        }
        println!("{:?}", dp);
        dp[0][l-1]
    }
}

fn main() {
    let ans = Solution::check_valid_string(String::from("(*)"));
    println!("{}", ans);
}