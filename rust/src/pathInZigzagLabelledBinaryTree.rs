// Path In Zigzag Labelled Binary Tree
// leetcode: https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
struct Solution {}
impl Solution {
    pub fn path_in_zig_zag_tree(label: i32) -> Vec<i32> {
        let (mut temp, mut n) = (label, 0);
        let mut label = label;
        let mut ans = vec![];
        while temp > 0 {
            temp >>= 1;
            n += 1;
        }
        if n % 2 == 0 {
            label = (1<<n) + (1<<(n-1)) - 1 - label;
        }
        while n > 0 {
            if n % 2 == 0 {
                ans.push((1<<n) + (1<<(n-1)) - 1 - label);
            } else {
                ans.push(label);
            }
            label >>= 1;
            n -= 1;
        }
        ans.reverse();
        ans
    }
}

fn main() {
    let ans = Solution::path_in_zig_zag_tree(14);
    println!("{:?}", ans);
}