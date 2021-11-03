// Unique Paths II
// leetcode: https://leetcode-cn.com/problems/unique-paths-ii/
struct Solution {}

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let n = obstacle_grid[0].len();
        let m = obstacle_grid.len();
        let mut dp = vec![vec![0; n]; m];
        for i in 0..n {
            if obstacle_grid[0][i] == 1 {
                break;
            }
            dp[0][i] = 1;
        }
        for i in 0..m {
            if obstacle_grid[i][0] == 1 {
                break;
            }
            dp[i][0] = 1;
        }
        for i in 1..m {
            for j in 1..n {
                if obstacle_grid[i][j] == 1 {
                    dp[i][j] = 0;
                    continue;
                }
                dp[i][j] = dp[i-1][j] + dp[i][j-1];
            }
        }
        dp[m-1][n-1]
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![0, 0]);
    let ans = Solution::unique_paths_with_obstacles(v);
    println!("{}", ans);
}