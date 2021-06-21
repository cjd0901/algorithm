// Search a 2D Matrix
// leetcode: https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
struct Solution {}

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut row = matrix.len() - 1;
        let mut r: i32 = row as i32;
        let mut col = 0;
        while r >= 0 && col < matrix[0].len() {
            if matrix[row][col] > target {
                r -= 1;
                if row > 0 {
                    row -= 1;
                }
            }else if matrix[row][col] < target {
                col += 1;
            }else {
                return true;
            }
        }
        false
    }
}

fn main() {
    let mut v = Vec::new();
    v.push(vec![-5]);
    let ans = Solution::search_matrix(v, -10);
    println!("{}", ans);
}