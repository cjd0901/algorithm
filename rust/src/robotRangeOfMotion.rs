// 机器人的运动范围
// leetcode: https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
struct Solution {}
impl Solution {
    pub fn moving_count(m: i32, n: i32, k: i32) -> i32 {
        let mut visited = vec![vec![false; n as usize]; m as usize];
        dfs(0, 0, m, n, k, &mut visited)
    }
}

fn dfs(i: i32, j: i32, m: i32, n: i32, k: i32, visited: &mut Vec<Vec<bool>>) -> i32 {
    if i < 0 || j < 0 || i >= m || j >= n || sums(i) + sums(j) > k || visited[i as usize][j as usize] {
        return 0
    }
    visited[i as usize][j as usize] = true;
    1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited) + dfs(i-1, j, m, n, k, visited) + dfs(i, j-1, m, n, k, visited)
}

fn sums(x: i32) -> i32 {
    let mut sum = 0;
    let mut x = x;
    while x != 0 {
        sum += x % 10;
        x = x / 10;
    }
    sum
}

fn main() {
    let ans = Solution::moving_count(1, 3, 11);
    println!("{}", ans);
}