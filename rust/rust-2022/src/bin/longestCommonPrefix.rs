// Longest Common Prefix
// leetcode: https://leetcode-cn.com/problems/longest-common-prefix/
struct Solution {}

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let chs = strs[0].chars().collect::<Vec<_>>();
        let mut max_length = chs.len();
        for i in 1..strs.len() {
            let shs = strs[i].chars().collect::<Vec<_>>();
            let l = chs.len().min(shs.len());
            let mut pre = 0;
            for j in 0..l {
                if chs[j] == shs[j] {
                    pre += 1;
                } else {
                    break;
                }
            }
            max_length = max_length.min(pre);
        }
        if max_length == 0 {
            return "".to_string();
        }
        chs[0..max_length].iter().collect()
    }
}

fn main() {
    let mut v = vec![];
    v.push(String::from("flower"));
    v.push(String::from("flow"));
    v.push(String::from("flight"));
    let ans = Solution::longest_common_prefix(v);
    println!("{}", ans);
}