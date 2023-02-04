// Combination Sum
// leetcode: https://leetcode-cn.com/problems/combination-sum/
struct Solution {}

impl Solution {
    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = Vec::new();
        let mut combile: Vec<i32> = Vec::new();
        dfs(&candidates, target, &mut combile, &mut ans, 0);
        ans
    }
}

fn dfs(candidates: &Vec<i32>, target: i32, combine: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>, idx: usize) {
    if idx == candidates.len() {
        return;
    }
    if target == 0 {
        ans.push(combine.to_vec());
        return;
    }
    dfs(candidates, target, combine, ans, idx+1);
    if target - candidates[idx] >= 0 {
        combine.push(candidates[idx]);
        dfs(candidates, target - candidates[idx], combine, ans, idx);
        combine.pop();
    }
}

fn main() {
    let ans = Solution::combination_sum(vec![2,3,6,7], 7);
    println!("{:?}", ans);
}