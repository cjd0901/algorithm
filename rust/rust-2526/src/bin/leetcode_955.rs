struct Solution {}

impl Solution {
    pub fn min_deletion_size(strs: Vec<String>) -> i32 {
        let (n, m) = (strs.len(), strs[0].len());
        let mut ans = 0;
        let mut cur = vec![String::new(); n];
        'next: for j in 0..m {
            for i in 0..n - 1 {
                let (mut a, mut b) = (cur[i].clone(), cur[i+1].clone());
                a.push(strs[i].chars().nth(j).unwrap());
                b.push(strs[i+1].chars().nth(j).unwrap());
                if a > b {
                    ans += 1;
                    continue 'next;
                }
            }
            for i in 0..n {
                cur[i].push(strs[i].chars().nth(j).unwrap());
            }
        }
        ans
    }
}

fn main() {}

