struct Solution {}

impl Solution {
    pub fn has_same_digits(s: String) -> bool {
        let mut chs: Vec<u32> = s.chars()
            .map(|ch| ch.to_digit(10).unwrap())
            .collect();
        let mut n = chs.len();
        while n > 2 {
            for i in 0..n-1 {
                chs[i] = (chs[i] + chs[i+1]) % 10;
            }
            n -= 1;
        }

        chs[0] == chs[1]
    }
}

fn main() {}