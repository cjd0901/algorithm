struct Solution {}

impl Solution {
    pub fn maximum_happiness_sum(mut happiness: Vec<i32>, k: i32) -> i64 {
        let mut ans = 0i64;
        happiness.sort_by(|a, b| b.cmp(a));
        let mut pick_count = 0;
        for &h in happiness.iter() {
            if pick_count == k {
                break;
            }
            ans += (h - pick_count).max(0) as i64;
            pick_count += 1;
        }
        ans
    }
}

fn main() {}