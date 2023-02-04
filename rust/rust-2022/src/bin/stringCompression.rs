// String Compression
// leetcode: https://leetcode-cn.com/problems/string-compression/
struct Solution {}
impl Solution {
    pub fn compress(chars: &mut Vec<char>) -> i32 {
        let (mut pre, mut count) = (chars[0], 1);
        let mut idx = 0;
        for i in 1..chars.len() {
            if pre == chars[i] {
                count += 1;
            } else {
                chars[idx] = pre;
                idx += 1;
                if count > 1 {
                    for ch in count.to_string().chars() {
                        chars[idx] = ch;
                        idx += 1;
                    }
                }
                count = 1;
                pre = chars[i];
            }
        }
        chars[idx] = pre;
        idx += 1;
        if count > 1 {
            for ch in count.to_string().chars() {
                chars[idx] = ch;
                idx += 1;
            }
        }
        idx as i32
    }
}

fn main() {
    let ans = Solution::compress(&mut vec!['a', 'b', 'b', 'c', 'c', 'c']);
    println!("{}", ans);
}