// Friends Of Appropriate Ages
// leetcode: https://leetcode-cn.com/problems/friends-of-appropriate-ages/
struct Solution {}

impl Solution {
    pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
        let mut ages = ages;
        let mut ans = 0;
        ages.sort();
        let (mut left, mut right) = (0, 0);
        for age in ages.iter() {
            if *age < 15 {
                continue;
            }

            while (ages[left as usize] * 2) <= *age + 14 {
                left += 1;
            }

            while ((right + 1) as usize) < ages.len() && ages[right+1] <= *age {
                right += 1;
            }
            ans += (right - left) as i32;
        }
        ans
    }
}

fn main() {
    let ans = Solution::num_friend_requests(vec![101,56,69,48,30]);
    println!("{}", ans);
}