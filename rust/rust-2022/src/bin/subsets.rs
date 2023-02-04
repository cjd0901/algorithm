// Subsets
// leetcode: https://leetcode-cn.com/problems/subsets/
struct Solution {}

impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = vec![];
        let mut combine: Vec<i32> = Vec::new();
        dfs(&nums, &mut ans, &mut combine, 0);
        ans
    }
}

fn dfs(nums: &Vec<i32>, ans: &mut Vec<Vec<i32>>, combine: &mut Vec<i32>, cur: usize) {
    if cur == nums.len() {
        ans.push(combine.to_vec());
        return;
    }
    combine.push(nums[cur]);
    dfs(nums, ans, combine, cur+1);
    combine.pop();
    dfs(nums, ans, combine, cur+1);
}

fn main() {
    let ans = Solution::subsets(vec![1, 2, 3]);
    println!("{:?}", ans);
}