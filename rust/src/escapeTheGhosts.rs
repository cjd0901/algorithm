// Escape The Ghosts
// leetcode: https://leetcode-cn.com/problems/escape-the-ghosts/
struct Solution {}
impl Solution {
    pub fn escape_ghosts(ghosts: Vec<Vec<i32>>, target: Vec<i32>) -> bool {
        let dis = target[0].abs()+target[1].abs();
        for ghost in ghosts.iter() {
            let d = (ghost[0]-target[0]).abs() + (ghost[1]-target[1]).abs();
            if d <= dis {
                return false
            }
        }
        true
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![1,0]);
    let ans = Solution::escape_ghosts(v, vec![2, 0]);
    println!("{}", ans);
}