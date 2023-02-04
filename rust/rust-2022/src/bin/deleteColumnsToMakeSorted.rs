// Delete Columns to Make Sorted
// leetcode: https://leetcode-cn.com/problems/delete-columns-to-make-sorted/
struct Solution {}
impl Solution {
    pub fn min_deletion_size(strs: Vec<String>) -> i32 {
        let mut count = 0;
        let mut byte_str = vec![];
        for s in strs {
            let temp = s.chars().collect::<Vec<_>>();
            byte_str.push(temp);
        }
        for i in 0..byte_str[0].len() {
            for j in 0..byte_str.len()-1 {
                if byte_str[j+1][i] < byte_str[j][i] {
                    count += 1;
                    break;
                }
            }
        }
        count
    }
}

fn main() {
    let count = Solution::min_deletion_size(vec![String::from("zyx"), String::from("wvu"), String::from("tsr")]);
    println!("{}", count);
}