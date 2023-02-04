// Maximal Square
// leetcode: https://leetcode-cn.com/problems/maximal-square/
struct Solution {}
impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        let (rows, cols) = (matrix.len(), matrix[0].len());
        let mut dp = vec![vec![0; cols]; rows];
        let mut max = 0;
        for i in 0..rows {
            for j in 0..cols {
                if matrix[i][j] == 1 {
                    if j == 0 || i == 0 {
                        dp[i][j] = 1;
                    }else {
                        dp[i][j] = dp[i-1][j].min(dp[i][j-1]).min(dp[i-1][j-1])+1;
                    }
                    max += dp[i][j]
                }
            }
        }
        max
    }
}

fn main() {
    let mut matrix = Vec::new();
    matrix.push(vec![0, 1]);
    matrix.push(vec![1, 0]);
    let ans = Solution::maximal_square(matrix);
    println!("{}", ans);
}