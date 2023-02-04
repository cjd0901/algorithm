// Count Negative Numbers in a Sorted Matrix
// leetcode: https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
struct Solution {}
impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = ((grid.len() - 1) as i32, (grid[0].len() - 1) as i32);
        let mut ans = 0;
        let mut f = vec![vec![false; (n+1) as usize]; (m+1) as usize];
        dfs(&grid, &mut f, m, n, &mut ans);
        ans
    }
}

pub fn dfs(grid: &Vec<Vec<i32>>, f: &mut Vec<Vec<bool>>, row: i32, col: i32, ans: &mut i32) {
    if (row >= 0 && col >= 0) && grid[row as usize][col as usize] < 0 && !f[row as usize][col as usize] {
        *ans += 1;
        f[row as usize][col as usize] = true;
        dfs(grid, f, row, col-1, ans);
        dfs(grid, f, row-1, col, ans);
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![4,3,2,-1]);
    v.push(vec![3,2,1,-1]);
    v.push(vec![1,1,-1,-2]);
    v.push(vec![-1,-1,-2,-3]);
    let ans = Solution::count_negatives(v);
    println!("{}", ans);
}