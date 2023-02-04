// Smallest k
// leetcode: https://leetcode-cn.com/problems/smallest-k-lcci/
struct Solution {}
impl Solution {
    pub fn smallest_k(arr: Vec<i32>, k: i32) -> Vec<i32> {
        let mut arr = arr;
        arr.sort();
        arr[0..k as usize].to_vec()
    }
}

fn main() {
    let ans = Solution::smallest_k(vec![1,3,5,7,2,4,6,8], 4);
    println!("{:?}", ans)
}