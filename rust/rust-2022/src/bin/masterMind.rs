// Master Mind
// leetcode: https://leetcode-cn.com/problems/master-mind-lcci/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn master_mind(solution: String, guess: String) -> Vec<i32> {
        let sol: Vec<u8> = solution.bytes().collect();
        let gue: Vec<u8> = guess.bytes().collect();
        let mut map = HashMap::new();
        let mut ans = vec![0; 2];
        for s in &sol {
            let count = map.entry(s).or_insert(0);
            *count += 1;
        }
        for i in 0..4 {
            if sol[i] == gue[i] {
                ans[0] += 1;
                if let Some(count) = map.get_mut(&gue[i]) {
                    if *count > 0 {
                        *count -= 1;
                    }
                }
            }
        }
        for i in 0..4 {
            if sol[i] != gue[i] {
                if let Some(count) = map.get_mut(&gue[i]) {
                    if *count > 0 {
                        *count -= 1;
                        ans[1] += 1;
                    }
                }
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::master_mind(String::from("RGBY"), String::from("GGRR"));
    println!("{:?}", ans);
}