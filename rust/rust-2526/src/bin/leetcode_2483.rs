struct Solution {}

impl Solution {
    pub fn best_closing_time(customers: String) -> i32 {
        let mut ans = 0;
        let mut cost = customers.chars().filter(|&c| c == 'Y').count() as i32;
        let mut prev = cost;
        for (i, c) in customers.chars().enumerate() {
            match c {
                'Y' => cost -= 1,
                'N' => cost += 1,
                _ => ()
            }
            if cost < prev {
                ans = i as i32 + 1;
                prev = cost;
            }
        }
       ans
    }
}

fn main() {
    Solution::best_closing_time(String::from("YYNY"));
}