struct Solution {}

//2373. 矩阵中的局部最大值
impl Solution {
    pub fn largest_local(grid: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let len = grid.len();
        let mut max_local = vec![vec![0; len - 2]; len - 2];
        for i in 0..len {
            for j in 0..len {
                for x in 0..len - 2 {
                    for y in 0..len - 2 {
                        let (a, b) = (x + 1, y + 1);
                        if i >= a - 1 && i <= a + 1 && j >= b - 1 && j <= b + 1 {
                            max_local[x][y] = max_local[x][y].max(grid[i][j]);
                        }
                    }
                }
            }
        }
        max_local
    }
}

fn main() {}
