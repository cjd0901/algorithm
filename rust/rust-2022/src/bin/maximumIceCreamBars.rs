// Maximum Ice Cream Bars
// leetcode: https://leetcode-cn.com/problems/maximum-ice-cream-bars/
struct Solution {}
impl Solution {
    pub fn max_ice_cream(costs: Vec<i32>, coins: i32) -> i32 {
        let mut costs = costs;
        let mut coins = coins;
        let mut max = 0;
        let mut i: usize = 0;
        costs.sort();
        while i < costs.len() && coins >= costs[i] {
            coins -= costs[i];
            max += 1;
            i += 1;
        }
        max
    }
}

fn main() {
    let ans = Solution::max_ice_cream(vec![1,3,2,4,1], 7);
    println!("{}", ans);
}