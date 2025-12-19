struct Solution {}

impl Solution {
    fn is_prime(num: i32) -> bool {
        match num {
            0 | 1 => false,
            2 => true,
            _ if num % 2 == 0 => false,
            _ => {
                (3..).step_by(2).
                    take_while(|&x| x * x <= num).
                    all(|x| num % x != 0)
            }
        }
    }

    pub fn maximum_prime_difference(nums: Vec<i32>) -> i32 {
        let (mut start, mut end) = (-1, -1);
        for (i, &num) in nums.iter().enumerate() {
            if Self::is_prime(num) {
                if start == -1 {
                    start = i as i32;
                }
                end = i as i32;
            }
        }
        end - start
    }
}

fn main() {

}