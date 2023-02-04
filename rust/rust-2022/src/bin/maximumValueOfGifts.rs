// Maximum value of gifts
// leetcode: https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
struct Solution {}
impl Solution {
    pub fn max_value(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        let mut dp = vec![vec![0; n+1]; m+1];
        for i in 1..m+1 {
            for j in 1..n+1 {
                let num = grid[i-1][j-1];
                dp[i][j] = (num+dp[i-1][j]).max(num+dp[i][j-1]);
            }
        }
        dp[m][n]
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![1,3,1]);
    v.push(vec![1,5,1]);
    v.push(vec![4,2,1]);
    let ans = Solution::max_value(v);
    println!("{}", ans);
}