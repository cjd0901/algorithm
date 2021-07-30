// Longest Happy String
// leetcode: https://leetcode-cn.com/problems/longest-happy-string/
struct Solution {}

#[derive(Debug)]
struct Happy {
    val: i32,
    ch: char,
}

impl Solution {
    pub fn longest_diverse_string(a: i32, b: i32, c: i32) -> String {
        let mut v = vec![];
        v.push(Happy{val: a, ch: 'a'});
        v.push(Happy{val: b, ch: 'b'});
        v.push(Happy{val: c, ch: 'c'});
        let mut chars = vec![];
        loop {
            let n = chars.len();
            v.sort_by(|x, y| x.val.cmp(&y.val));
            if n > 1 && chars[n-1] == v[2].ch && chars[n-2] == v[2].ch {
                if v[1].val > 0 {
                    chars.push(v[1].ch);
                    v[1].val -= 1;
                } else {
                    break
                }
            } else {
                if v[2].val > 0 {
                    chars.push(v[2].ch);
                    v[2].val -= 1;
                } else {
                    break
                }
            }
        }
        chars.iter().collect()
    }
}

fn main() {
    let ans = Solution::longest_diverse_string(1, 1, 7);
    println!("{}", ans);
}