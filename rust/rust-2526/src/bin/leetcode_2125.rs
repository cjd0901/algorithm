struct Solution {}

impl Solution {
    pub fn number_of_beams(bank: Vec<String>) -> i32 {
        let mut ans = 0;
        let mut count = 0;
        for s in bank.iter() {
            let temp = s.chars().filter(|&c| c == '1').count() as i32;
            if temp == 0 {
                continue;
            }
            ans += temp * count;
            count = temp;
        }
        ans
    }
}

fn main() {}