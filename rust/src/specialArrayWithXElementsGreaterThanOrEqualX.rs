// Special Array With X Elements Greater Than or Equal X
// leetcode: https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
struct Solution {}
impl Solution {
    pub fn special_array(nums: Vec<i32>) -> i32 {
        let mut max = 0;
        for num in nums.iter() {
            max = max.max(*num);
        }
        let mut arr = vec![0; (max+1) as usize];
        for num in nums.iter() {
            arr[*num as usize] += 1;
        }
        let mut sum = 0;
        for i in (0..max+1).rev() {
            sum += arr[i as usize];
            if sum == i {
                return i;
            }
        }
        -1
    }
}

fn main() {
    let ans = Solution::special_array(vec![3, 5]);
    println!("{}", ans);
}