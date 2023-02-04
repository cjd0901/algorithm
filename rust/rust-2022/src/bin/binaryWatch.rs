// Binary Watch
// https://leetcode-cn.com/problems/binary-watch/

struct Solution {}
impl Solution {
    pub fn read_binary_watch(turned_on: i32) -> Vec<String> {
        let mut v: Vec<String> = Vec::new();
        for i in 0..12 {
            for j in 0..60 {
                if bit_count(i) + bit_count(j) == turned_on {
                    v.push(String::from(format!("{}:{}", i, if j >= 10 { format!("{}", j) } else { format!("0{}", j) })))
                }
            }
        }
        v
    }
}

fn bit_count(x: i32) -> i32 {
    let mut count = 0;
    let mut n = x;
    while n != 0 {
        n = n&(n-1);
        count = count + 1;
    }
    count
}

fn main() {
    let ans = Solution::read_binary_watch(1);
    println!("{:?}", ans);
}