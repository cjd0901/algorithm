// String Rotation
// leetcode: https://leetcode-cn.com/problems/string-rotation-lcci/
struct Solution {}
impl Solution {
    pub fn is_fliped_string(s1: String, s2: String) -> bool {
        let mut s2 = s2;
        if s1.len() != s2.len() {
            return false;
        }
        s2.push_str(&s2.to_string());
        s2.contains(&s1)
    }
}

fn main() {
    let ans = Solution::is_fliped_string(String::from("waterbottle"), String::from("erbottlewat"));
    println!("{}", ans);
}