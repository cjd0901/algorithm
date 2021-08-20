// Water Bottles
// leetcode: https://leetcode-cn.com/problems/water-bottles/
struct Solution {}
impl Solution {
    pub fn num_water_bottles(num_bottles: i32, num_exchange: i32) -> i32 {
        let mut ans = num_bottles;
        let mut num = num_bottles;
        while num >= num_exchange {
            ans += num/num_exchange;
            let m = num % num_exchange;
            num = num/num_exchange + m;
        }
        ans
    }
}

fn main() {
    let ans = Solution::num_water_bottles(2, 3);
    println!("{}", ans);
}