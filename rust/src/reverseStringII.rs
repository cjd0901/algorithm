// Reverse String II
// leetcode: https://leetcode-cn.com/problems/reverse-string-ii/
struct Solution {}
impl Solution {
    pub fn reverse_str(s: String, k: i32) -> String {
        let k = k as usize;
        let mut chs: Vec<char> = s.chars().collect();
        let mut i = 0 as usize;
        while i < s.len() {
            let (mut s, mut e) = (i, 0);
            e = (i+k-1).min(chs.len()-1);
            while s < e {
                let temp = chs[e];
                chs[e] = chs[s];
                chs[s] = temp;
                s += 1;
                e -= 1;
            }
            i += 2*k;
        }
        chs.iter().collect()
    }
}

fn main() {
    let ans = Solution::reverse_str(String::from("abcdefg"), 2);
    println!("{}", ans);
}