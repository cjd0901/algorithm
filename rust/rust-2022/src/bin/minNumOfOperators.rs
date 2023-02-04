struct Solution {}

impl Solution {
    pub fn min_num_of_operators(s: String) -> i32 {
        let s = s.chars().collect::<Vec<_>>();
        let len = s.len();
        let (mut zero, mut one) = (0, 0);
        for n in s.iter() {
            if *n == '0' {
                zero += 1;
            } else {
                one += 1;
            }
        }
        let mut temp = one;
        let mut operates = 0;
        for i in 0..zero as usize {
            if s[i] == '1' {
                temp -= 1;
                operates += 1;
            }
        }
        operates += one - temp;
        zero.min(one).min(operates)
    }
}

fn main() {
    let ans = Solution::min_num_of_operators(String::from("11110000"));
    println!("{}", ans);
}