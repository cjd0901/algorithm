struct Solution {}

impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let n = grid.len();
        let m = grid[0].len();
        let mut ans = 0;
        for i in (0..n).rev() {
           for j in (0..m).rev() {
               if grid[i][j] >= 0 {
                   break;
               }
               ans += 1;
           }
        }
        ans
    }
}

fn main() {}