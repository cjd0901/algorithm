struct Solution {}

impl Solution {
    pub fn get_maximum_consecutive(coins: Vec<i32>) -> i32 {
        let mut coins = coins;
        coins.sort();
        let mut ans = 1;
        for it in coins.iter() {
            if *it > ans {
                break;
            }
            ans += *it;
        }
        ans
    }
}

fn main() {
    let input = vec![1, 3];
    let ans = Solution::get_maximum_consecutive(input);
    println!("{ans}");
}
