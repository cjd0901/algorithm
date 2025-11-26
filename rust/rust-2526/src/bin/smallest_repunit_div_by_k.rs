use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn smallest_repunit_div_by_k(k: i32) -> i32 {
        let mut res = 1;
        let mut x = 0;
        let mut set: HashSet<i32> = HashSet::new();
        loop {
            x += 1;
            let t = res % k;
            if set.get(&t).is_some() {
                return -1;
            }

            if t % k == 0 {
                break;
            } else {
                set.insert(t);
            }
            res = (t * 10 + 1) % k;
        }
        x
    }
}

fn main () {
    let ans = Solution::smallest_repunit_div_by_k(1);
    println!("{}", ans);
}