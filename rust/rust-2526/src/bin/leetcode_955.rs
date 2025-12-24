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

    pub fn min_deletion_size2(strs: Vec<String>) -> i32 {
        let (n, m) = (strs.len(), strs[0].len());
        let mut ans = 0;
        let mut is_sorted = vec![false; n-1];
        'next: for j in 0..m {
            let mut new_sorted = is_sorted.clone();
            for i in 1..n {
                let a = strs[i-1].chars().nth(j).unwrap();
                let b = strs[i].chars().nth(j).unwrap();
                if !is_sorted[i-1] {
                    if a > b {
                        ans += 1;
                        continue 'next;
                    } else if a < b {
                        new_sorted[i-1] = true;
                    }
                }
            }
            is_sorted = new_sorted;
        }
       ans
    }
}

fn main() {
    let _ = Solution::min_deletion_size2(
        vec![
            String::from("vdy"),
            String::from("vei"),
            String::from("zvc"),
            String::from("zld")
        ]
    );
}

