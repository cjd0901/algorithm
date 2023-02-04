// Valid Sudoku
// leetcode: https://leetcode-cn.com/problems/valid-sudoku/
struct Solution {}
impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let (mut row, mut col, mut grid) =
            (vec![vec![false; 9]; 9], vec![vec![false; 9]; 9], vec![vec![false; 9]; 9]);
        for i in 0..9 {
            for j in 0..9 {
                let ch = board[i][j];
                if ch != '.' {
                    let idx = (ch as u8 - '0' as u8 - 1) as usize;
                    let grid_idx = ((i / 3) * 3 + j / 3) as usize;
                    if row[i][idx] || col[j][idx] || grid[grid_idx][idx] {
                        return false;
                    }
                    row[i][idx] = true;
                    col[j][idx] = true;
                    grid[grid_idx][idx] = true;
                }
            }
        }
        true
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec!['5','3','.','.','7','.','.','.','.']);
    v.push(vec!['6','.','.','1','9','5','.','.','.']);
    v.push(vec!['.','9','8','.','.','.','.','6','.']);
    v.push(vec!['8','.','.','.','6','.','.','.','3']);
    v.push(vec!['4','.','.','8','.','3','.','.','1']);
    v.push(vec!['7','.','.','.','2','.','.','.','6']);
    v.push(vec!['.','6','.','.','.','.','2','8','.']);
    v.push(vec!['.','.','.','4','1','9','.','.','5']);
    v.push(vec!['.','.','.','.','8','.','.','7','9']);
    let ans = Solution::is_valid_sudoku(v);
    println!("{}", ans);
}