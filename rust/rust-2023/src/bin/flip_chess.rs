struct Solution {}

use std::collections::VecDeque;

const DIRS: [(i32, i32); 8] = [
    (1, 0),
    (-1, 0),
    (0, -1),
    (0, 1),
    (1, -1),
    (1, 1),
    (-1, 1),
    (-1, -1),
];

impl Solution {
    pub fn flip_chess(chessboard: Vec<String>) -> i32 {
        fn bfs(mut chessboard: Vec<Vec<char>>, cur: (usize, usize)) -> i32 {
            let (m, n) = (chessboard.len(), chessboard[0].len());
            let mut q = VecDeque::new();
            q.push_back(cur);
            chessboard[cur.0][cur.1] = 'X';
            let mut cnt = 0;
            while !q.is_empty() {
                let cur = q.pop_front().unwrap();
                for dir in DIRS.iter() {
                    let (mut n_x, mut n_y) = (cur.0 as i32, cur.1 as i32);
                    loop {
                        n_x += dir.0;
                        n_y += dir.1;
                        if n_x < 0 || n_x >= m as i32 || n_y < 0 || n_y >= n as i32 {
                            break;
                        }
                        match chessboard[n_x as usize][n_y as usize] {
                            'X' => {
                                n_x -= dir.0;
                                n_y -= dir.1;
                                while n_x != cur.0 as i32 || n_y != cur.1 as i32 {
                                    chessboard[n_x as usize][n_y as usize] = 'X';
                                    q.push_back((n_x as usize, n_y as usize));
                                    cnt += 1;
                                    n_x -= dir.0;
                                    n_y -= dir.1;
                                }
                                break;
                            }
                            'O' => {
                                continue;
                            }
                            _ => {
                                break;
                            }
                        }
                    }
                }
            }
            cnt
        }
        let chessboard: Vec<Vec<char>> = chessboard
            .into_iter()
            .map(|s| s.chars().collect())
            .collect();
        let mut ans = 0;
        for x in 0..chessboard.len() {
            for y in 0..chessboard[0].len() {
                if chessboard[x][y] == '.' {
                    ans = ans.max(bfs(chessboard.clone(), (x, y)));
                }
            }
        }
        ans
    }
}

fn main() {}
