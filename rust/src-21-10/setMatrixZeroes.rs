// Set Matrix Zeroes
// leetcode: https://leetcode-cn.com/problems/set-matrix-zeroes/
struct Solution {}

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let m = matrix.len();
        let n = matrix[0].len();
        let (mut row, mut col) = (false, false);
        for i in 0..m {
            if matrix[i][0] == 0 {
                col = true;
                break;
            }
        }
        for i in 0..n {
            if matrix[0][i] == 0 {
                row = true;
                break;
            }
        }
        for i in 1..m {
            for j in 1..n {
                if matrix[i][j] == 0 {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }
        for i in 1..m {
            for j in 1..n {
                if matrix[i][0] == 0 || matrix[0][j] == 0 {
                    matrix[i][j] = 0;
                }
            }
        }
        if row {
            for i in 0..n {
                matrix[0][i] = 0;
            }
        }
        if col {
            for i in 0..m {
                matrix[i][0] = 0;
            }
        }
    }
}

fn main() {
    let mut matrix: Vec<Vec<i32>> = Vec::new();
    matrix.push(vec![0, 1, 2, 0]);
    matrix.push(vec![3, 4, 5, 2]);
    matrix.push(vec![1, 3, 1, 5]);
    Solution::set_zeroes(&mut matrix);
    println!("{:?}", matrix);
}