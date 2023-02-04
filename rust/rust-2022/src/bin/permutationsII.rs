// Permutations II
// leetcode: https://leetcode-cn.com/problems/permutations-ii/
struct Solution {}

impl Solution {
    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = Vec::new();
        let mut used = vec![false; nums.len()];
        let mut combine: Vec<i32> = Vec::new();
        let mut nums = nums;
        nums.sort();
        dfs(&nums, &mut ans, &mut combine, &mut used, 0);
        ans
    }
}

pub fn dfs(nums: &Vec<i32>, ans: &mut Vec<Vec<i32>>, combine: &mut Vec<i32>, used: &mut Vec<bool>, idx: usize) {
    if idx == nums.len() {
        ans.push(combine.to_vec());
        return;
    }
    for i in 0..nums.len() {
        if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue;
        }
        if !used[i] {
            used[i] = true;
            combine.push(nums[i]);
            println!("回溯前{:?}", combine);
            dfs(nums, ans, combine, used, idx+1);
            println!("回溯后{:?}", combine);
            used[i] = false;
            combine.pop();
        }
    }
}

fn main() {
    let ans = Solution::permute_unique(vec![1,1,2]);
    println!("{:?}", ans);
}