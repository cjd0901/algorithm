// Latest Time by Replacing Hidden Digits
// leetcode: https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits/
struct Solution {}
impl Solution {
    pub fn maximum_time(time: String) -> String {
        let mut ch:Vec<char> = time.chars().collect();
        if ch[0] == '?' {
            if ch[1] >= '4' && ch[1] <= '9' {
                ch[0] = '1';
            }else {
                ch[0] = '2';
            }
        }
        if ch[1] == '?' {
            if ch[0] == '2' {
                ch[1] = '3';
            }else {
                ch[1] = '9';
            }
        }
        if ch[3] == '?' {
            ch[3] = '5';
        }
        if ch[4] == '?' {
            ch[4] = '9';
        }
        ch.iter().collect()
    }
}

fn main() {
    let ans = Solution::maximum_time(String::from("2?:?0"));
    println!("{}", ans);
}