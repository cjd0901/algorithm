use std::str::FromStr;

struct Solution {}

// 2409. 统计共同度过的日子数
impl Solution {
    pub fn count_days_together(
        arrive_alice: String,
        leave_alice: String,
        arrive_bob: String,
        leave_bob: String,
    ) -> i32 {
        let month_days: [i32; 12] = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        let get_days = |x: String| -> i32 {
            let d_vec = x
                .as_str()
                .split("-")
                .map(|d| i32::from_str(d).unwrap())
                .collect::<Vec<i32>>();
            (0..(d_vec[0] - 1) as usize)
                .into_iter()
                .fold(0, |days, i| days + month_days[i])
                + d_vec[1]
        };
        let arrive_day = std::cmp::max(get_days(arrive_alice), get_days(arrive_bob));
        let leave_day = std::cmp::min(get_days(leave_alice), get_days(leave_bob));
        if arrive_day > leave_day {
            0
        } else {
            leave_day - arrive_day + 1
        }
    }
}

fn main() {}
