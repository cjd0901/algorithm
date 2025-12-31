use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn avoid_flood(rains: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![1; rains.len()];
        let mut map = HashMap::new();
        let mut sun = vec![];
        for (i, &rain) in rains.iter().enumerate() {
            if rain == 0 {
                sun.push(i);
            } else {
                ans[i] = -1;
                if let Some(day) = map.get(&rain) {
                    if let Err(j) = sun.binary_search(day) {
                        if j == sun.len() {
                            return vec![];
                        }

                        ans[sun[j]] = rain;
                        sun.remove(j);
                    }
                }
                map.insert(rain, i);
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::avoid_flood(vec![1,2,0,0,2,1]);
    println!("{:?}", ans);
}