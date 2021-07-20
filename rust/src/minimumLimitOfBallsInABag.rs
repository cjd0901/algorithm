// Minimum Limit of Balls in a Bag
// leetcode: https://leetcode-cn.com/problems/minimum-limit-of-balls-in-a-bag/
struct Solution {}
impl Solution {
    pub fn minimum_size(nums: Vec<i32>, max_operations: i32) -> i32 {
        let (mut l, mut r) = (1, 0);
        let mut ans = 0;
        for n in nums.iter() {
            if *n > r {
                r = *n;
            }
        }
        while l <= r {
            let mid = l + (r-l)/2;
            let mut opts = 0;
            for n in nums.iter() {
                opts += (n - 1) / mid;
            }
            if opts <= max_operations {
                ans = mid;
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::minimum_size(vec![1], 1);
    println!("{}", ans);
}