// Sub Sort
// leetcode: https://leetcode-cn.com/problems/sub-sort-lcci/
struct Solution {}
impl Solution {
    pub fn sub_sort(array: Vec<i32>) -> Vec<i32> {
        let mut copy = array.to_vec();
        copy.sort();
        let n = copy.len();
        let (mut left, mut right) = (0, (copy.len()-1) as i32);
        for i in 0..n {
            if array[i] != copy[i] {
                break
            }
            left += 1;
        }
        for i in (0..n).rev() {
            if array[i] != copy[i] {
                break
            }
            right -= 1;
        }
        if right == -1 {
            return vec![-1, -1];
        }
        vec![left, right]
    }
}

fn main() {
    let ans = Solution::sub_sort(vec![1,2,3,4]);
    println!("{:?}", ans);
}