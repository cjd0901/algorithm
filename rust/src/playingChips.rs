// Playing Chips
// leetcode: https://leetcode-cn.com/problems/minimum-cost-to-move-chips-to-the-same-position/
struct Solution {}
impl Solution {
    pub fn min_cost_to_move_chips(position: Vec<i32>) -> i32 {
        let (mut odd, mut even) = (0, 0);
        for pos in position.iter() {
            if *pos&1 == 1 {
                odd += 1;
            } else {
                even += 1;
            }
        }
        odd.min(even)
    }
}

fn main() {
    let ans = Solution::min_cost_to_move_chips(vec![2,2,2,3,3]);
    println!("{}", ans);
}