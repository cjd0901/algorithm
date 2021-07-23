// Jump Game
// leetcode: https://leetcode-cn.com/problems/jump-game/
struct Solution {}
impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut furthest = 0;
        for i in 0..nums.len() {
            if i as i32 <= furthest {
                furthest = furthest.max(i as i32 + nums[i]);
                if furthest >= nums.len() as i32-1 {
                    return true;
                }
            }
        }
        false
    }
}

fn main() {
    let ans = Solution::can_jump(vec![3,2,1,0,4]);
    println!("{}", ans);
}