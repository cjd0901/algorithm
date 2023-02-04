// Stone Game Center
// leetcode: https://leetcode-cn.com/problems/stone-game/
struct Solution {}

impl Solution {
    // pub fn stone_game(piles: Vec<i32>) -> bool {
    //     true
    // }
    pub fn stone_game(piles: Vec<i32>) -> bool {
        let n = piles.len();
        let mut dp: Vec<Vec<i32>> = vec![vec![0; n]; n];
        for i in 0..n {
            dp[i][i] = piles[i];
        }
        for i in (0..n-1).rev() {
            for j in i+1..n {
                dp[i][j] = (piles[i] - dp[i+1][j]).max(piles[j] - dp[i][j-1]);
            }
        }
        return dp[0][n-1] > 0
    }
}

fn main() {
    let ans:bool = Solution::stone_game(vec![5,3,4,5]);
    println!("{}", ans);
}
// 亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。

// 游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。

// 亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。

// 假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。
// [5, 3, 4, 5]