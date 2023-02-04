// H-index II
// leetcode: https://leetcode-cn.com/problems/h-index-ii/
struct Solution {}

impl Solution {
    pub fn h_index(citations: Vec<i32>) -> i32 {
        let n = citations.len();
        let (mut lo, mut hi) = (0, (n - 1) as i32);
        while lo <= hi {
            let mid = (lo + (hi - lo) / 2) as usize;
            if (n - mid) as i32 == citations[mid] {
                return (n - mid) as i32;
            }else if (n - mid) as i32 > citations[mid] {
                lo = mid as i32 + 1;
            }else {
                hi = mid as i32 - 1;
            }
        }
        n as i32 - lo
    }
}

fn main() {
    let ans = Solution::h_index(vec![100]);
    println!("{}", ans);
}