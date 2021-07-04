// Set Mismatch
// leetcode: https://leetcode-cn.com/problems/set-mismatch/
use::std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_error_nums(nums: Vec<i32>) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut ans = vec![0; 2];
        for n in &nums {
            let count = map.entry(*n).or_insert(0);
            *count += 1;
        }
        for i in 1..nums.len()+1 {
            let i = i as i32;
            if let Some(count) = map.get(&i) {
                if *count == 2 {
                    ans[0] = i as i32;
                }
            }else {
                ans[1] = i as i32;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::find_error_nums(vec![2,2]);
    println!("{:?}", ans);
}