// Count And Say
// leetcode: https://leetcode-cn.com/problems/count-and-say/
struct Solution {}

impl Solution {
    pub fn count_and_say(n: i32) -> String {
        let mut prev = "1".to_string();
        for _ in 1..n as usize {
            let mut sb = String::new();
            let chs = prev.chars().collect::<Vec<_>>();
            let (mut pos, mut start) = (0 as usize, 0 as usize);
            while pos < chs.len() {
                while pos < prev.len() && chs[pos] == chs[start] {
                    pos += 1;
                }
                // sb.push_str((pos-start).to_string().as_str());
                sb.push_str(&(pos-start).to_string());
                sb.push(chs[start]);
                start = pos;
            }
            prev = sb;
        }
        prev
    }
}

fn main() {
    let ans = Solution::count_and_say(3);
    println!("{}", ans);
}