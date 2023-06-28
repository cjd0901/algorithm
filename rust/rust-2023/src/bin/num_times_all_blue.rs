struct Solution {}

impl Solution {
    pub fn num_times_all_blue(flips: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut cur_max = -1;
        for i in 0..flips.len() {
            let flip = flips[i];
            cur_max = cur_max.max(flip);
            if cur_max == i as i32 + 1 {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {}
