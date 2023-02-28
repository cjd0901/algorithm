struct Solution {}

//1140.石子游戏II
impl Solution {
    pub fn stone_game_ii(piles: Vec<i32>) -> i32 {
        let mut len = piles.len();
        let mut dp = vec![vec![0; len + 1]; len];
        let mut sum = 0;
        for (i, pile) in piles.iter().enumerate().rev() {
            sum += *pile;
            for M in 1..len + 1 {
                if i + M * 2 >= len {
                    dp[i][M] = sum;
                } else {
                    for x in 1..M * 2 + 1 {
                        dp[i][M] = dp[i][M].max(sum - dp[i + x][M.max(x)]);
                    }
                }
            }
        }
        dp[0][1]
    }
}

fn main() {}
