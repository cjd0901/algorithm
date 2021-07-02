// Partition Array Into Three Parts With Equal Sum
// leetcode: https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
struct Solution {}
impl Solution {
    pub fn can_three_parts_equal_sum(arr: Vec<i32>) -> bool {
        let mut pre_sum = vec![0; arr.len()];
        pre_sum[0] = arr[0];
        for i in 1..arr.len() {
            let i = i as usize;
            pre_sum[i] = pre_sum[i-1] + arr[i];
        }
        let (mut first, mut second) = (false, false);
        let last = pre_sum[arr.len()-1];
        if last % 3 != 0 {
            return false
        }
        for i in 0..pre_sum.len() {
            let n = pre_sum[i as usize];
            if !first && n == last / 3 {
                first = true;
            } else if first && i != pre_sum.len()-1 && !second && n == last / 3 * 2 {
                second = true;
            }
        }
        first && second
    }
}

fn main() {
    let ans = Solution::can_three_parts_equal_sum(vec![1, -1, 1, -1]);
    println!("{}", ans);
}