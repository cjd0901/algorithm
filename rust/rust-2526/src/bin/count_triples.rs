struct Solution {}

impl Solution {
    pub fn count_triples(n: i32) -> i32 {
        let mut ans = 0;
        for a in 1..=n {
            for b in 1..=n {
                let x = a * a + b * b;
                let c = (x as f64).sqrt() as i32;
                if c * c == x && c <= n {
                    ans += 1;
                }
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::count_triples(10);
    println!("{}", ans);
}