// Mirror Reflection
// leetcode: https://leetcode-cn.com/problems/mirror-reflection/
struct Solution {}
impl Solution {
    pub fn mirror_reflection(p: i32, q: i32) -> i32 {
        let (mut p, mut q) = (p, q);
        while p&1 == 0 && q&1 == 0 {
            p >>= 1;
            q >>= 1;
        }
        if p&1 == 0 {
            return 2;
        }
        if q&1 == 0 {
            return 0;
        }
        1
    }
}

fn main() {
    let ans = Solution::mirror_reflection(2, 1);
    println!("{}", ans);
}