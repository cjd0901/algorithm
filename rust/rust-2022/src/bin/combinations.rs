// Combinations
// leetcode: https://leetcode-cn.com/problems/combinations/
struct Solution {}

impl Solution {
    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = Vec::new();
        let mut combine: Vec<i32> = Vec::new();
        dfs(&mut ans, &mut combine, 1, n, k);
        ans
    }
}

fn dfs(ans: &mut Vec<Vec<i32>>, combine: &mut Vec<i32>, cur: i32, n: i32, k: i32) {
    if combine.len() as i32 + n - cur + 1 < k {
        return;
    }
    if combine.len() as i32 == k {
        ans.push(combine.to_vec());
        return;
    }
    if cur == n+1 {
        return;
    }
    combine.push(cur);
    dfs(ans, combine, cur + 1, n, k);
    combine.pop();
    dfs(ans, combine, cur + 1, n, k);
}

fn main() {
    let ans = Solution::combine(4, 2);
    println!("{:?}", ans);
}