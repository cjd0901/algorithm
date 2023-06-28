use std::collections::VecDeque;

struct Solution {}

const DIRS: [(i32, i32); 4] = [(0, 1), (0, -1), (1, 0), (-1, 0)];

impl Solution {
    pub fn closed_island(mut grid: Vec<Vec<i32>>) -> i32 {
        fn bfs(grid: &mut Vec<Vec<i32>>, start: (usize, usize)) -> bool {
            let (m, n) = (grid.len(), grid[0].len());
            let mut is_closed_island = true;
            let mut q = VecDeque::new();
            q.push_back(start);
            while !q.is_empty() {
                let cur = q.pop_front().unwrap();
                if cur.0 == 0 || cur.0 == m - 1 || cur.1 == 0 || cur.1 == n - 1 {
                    is_closed_island = false;
                }
                grid[cur.0][cur.1] = 1;
                for dir in DIRS.iter() {
                    let (n_x, n_y) = (cur.0 as i32 + dir.0, cur.1 as i32 + dir.1);
                    if n_x < 0 || n_x == m as i32 || n_y < 0 || n_y == n as i32 {
                        continue;
                    }
                    let (n_x, n_y) = (n_x as usize, n_y as usize);
                    if grid[n_x][n_y] == 0 {
                        q.push_back((n_x, n_y));
                    }
                }
            }
            is_closed_island
        }
        let mut ans = 0;
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == 0 && bfs(&mut grid, (i, j)) {
                    ans += 1;
                }
            }
        }
        ans
    }
}

fn main() {}
