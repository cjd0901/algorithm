struct Solution {}

impl Solution {
    pub fn first_day_been_in_all_rooms(next_visit: Vec<i32>) -> i32 {
        const MOD: i64 = 1_000_000_007;
        let n = next_visit.len();
        let mut s = vec![0; n];
        for i in 0..n - 1 {
            let j = next_visit[i] as usize;
            s[i + 1] = (s[i] * 2 - s[j] + 2 + MOD) % MOD;
        }
        s[n - 1] as _
    }
}

fn main() {
    let ans = Solution::first_day_been_in_all_rooms(vec![0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]);
    println!("{}",ans);
}