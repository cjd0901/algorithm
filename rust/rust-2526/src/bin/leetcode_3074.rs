struct Solution {}

impl Solution {
    pub fn minimum_boxes(apple: Vec<i32>, mut capacity: Vec<i32>) -> i32 {
        let mut remaining_apple = apple.iter().sum::<i32>();
        let mut ans = 0;
        capacity.sort();
        for i in (0..capacity.len()).rev() {
            if remaining_apple <= 0 {
                break;
            }
            remaining_apple -= capacity[i];

        }
        ans
    }
}

fn main() {}