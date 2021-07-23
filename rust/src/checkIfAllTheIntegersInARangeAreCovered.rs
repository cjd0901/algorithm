// Check if All the Integers in a Range Are Covered
// leetcode: https://leetcode-cn.com/problems/check-if-all-the-integers-in-a-range-are-covered/
struct Solution {}
impl Solution {
    pub fn is_covered(ranges: Vec<Vec<i32>>, left: i32, right: i32) -> bool {
        let mut diff = vec![0; 52];
        for i in 0..ranges.len() {
            let r = &ranges[i as usize];
            diff[r[0] as usize] += 1;
            diff[(r[1]+1) as usize] -= 1;
        }
        let mut cnt = 0;
        for i in 1..51 {
            cnt += diff[i];
            if cnt <= 0 && left <= i as i32 && i as i32 <= right {
                return false;
            }
        }
        true
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![1, 2]);
    v.push(vec![3, 4]);
    v.push(vec![5, 6]);
    let ans = Solution::is_covered(v, 2, 5);
    println!("{}", ans);
}