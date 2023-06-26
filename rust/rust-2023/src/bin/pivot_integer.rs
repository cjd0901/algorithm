struct Solution {}

impl Solution {
    pub fn pivot_integer(n: i32) -> i32 {
        let pre_sums = (1..=n)
            .into_iter()
            .fold(vec![0; n as usize + 1], |mut pre_sums, x| {
                pre_sums[x as usize] = pre_sums[x as usize - 1] + x;
                pre_sums
            });
        println!("{:?}", pre_sums);
        for i in 1..=n {
            if pre_sums[i as usize] - pre_sums[0] == pre_sums[n as usize] - pre_sums[i as usize - 1]
            {
                return i;
            }
        }
        -1
    }
}

fn main() {}
