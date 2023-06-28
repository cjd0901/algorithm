struct Solution {}

impl Solution {
    pub fn mice_and_cheese(reward1: Vec<i32>, reward2: Vec<i32>, mut k: i32) -> i32 {
        let mut rewards = reward1
            .into_iter()
            .zip(reward2.into_iter())
            .collect::<Vec<(i32, i32)>>();
        rewards.sort_by(|a, b| (b.0 - b.1).cmp(&(a.0 - a.1)));
        rewards.into_iter().fold(0, |ans, reward| {
            if k > 0 {
                k -= 1;
                ans + reward.0
            } else {
                ans + reward.1
            }
        })
    }
}

fn main() {}
