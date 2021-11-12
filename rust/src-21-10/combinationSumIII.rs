// Combination Sum III
// leetcode: https://leetcode-cn.com/problems/combination-sum-iii/
struct Solution {}

impl Solution {
    pub fn combination_sum3(k: i32, n: i32) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = Vec::new();
        let mut combine: Vec<i32> = Vec::new();
        dfs(&mut combine, &mut ans, k, 0, 0, n);
        ans
    }
}

fn dfs(combine: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>, k: i32, cur: i32, sum: i32, n: i32) {
    if sum > n || combine.len() as i32 > k {
        return
    }
    if sum == n {
        if combine.len() as i32 == k {
            ans.push(combine.to_vec());
        }
        return;
    }
    for i in cur+1..10 {
        let temp = sum + i;
        combine.push(i);
        dfs(combine, ans, k, i, temp, n);
        combine.pop();
    }
}

fn main() {
    let ans = Solution::combination_sum3(9, 45);
    println!("{:?}", ans);
}