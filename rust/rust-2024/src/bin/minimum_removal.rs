struct Solution {}

impl Solution {
    pub fn minimum_removal(mut beans: Vec<i32>) -> i64 {
        beans.sort();
        let total_beans: i64 = beans.iter().fold(0, |total, bean| total + *bean as i64);
        let count = beans.len();
        let mut ans = i64::MAX;
        for (i, bean) in beans.iter().enumerate() {
            ans = ans.min(total_beans - (count - i) as i64 * *bean as i64);
        }
        ans
    }
}

fn main () {}