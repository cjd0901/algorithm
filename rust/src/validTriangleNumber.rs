// Valid Triangle Number
// leetcode: https://leetcode-cn.com/problems/valid-triangle-number/
struct Solution {}
impl Solution {
    pub fn triangle_number(nums: Vec<i32>) -> i32 {
        let (mut nums, mut ans) = (nums, 0);
        let n = nums.len();
        nums.sort();
        let mut i = 0;
        while i+2 <= n {
            for j in i+1..n {
                let idx = match nums[j+1..].binary_search(&(nums[i]+nums[j]-1)) {
                    Ok(k) => {
                        k+1
                    },
                    Err(k) => {
                        k
                    }
                };
                ans += idx as i32;
            }
            i += 1;
        }
        ans
    }

    pub fn triangle_number2(nums: Vec<i32>) -> i32 {
        let (mut nums, mut ans) = (nums, 0);
        let n = nums.len();
        nums.sort();
        let mut i = 0;
        while i+2 <= n {
            for j in i+1..n-1 {
                let mut k = j+1;
                while k < n && nums[k] < nums[i] + nums[j]  {
                    ans += 1;
                    k += 1;
                }
            }
            i += 1;
        }
        ans as i32
    }
}

fn main() {
    let ans = Solution::triangle_number2(vec![4,2,3,4]);
    println!("{}", ans);
}