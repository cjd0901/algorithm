use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn equal_pairs(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        let mut map = HashMap::new();
        for i in 0..n {
            let mut s = String::new();
            for j in 0..m {
                s.push_str(grid[i][j].to_string().as_str());
                s.push('-');
            }
            map.entry(s).or_insert((0, 0)).0 += 1;
        }
        for j in 0..m {
            let mut s = String::new();
            for i in 0..m {
                s.push_str(grid[i][j].to_string().as_str());
                s.push('-');
            }
            map.entry(s).or_insert((0, 0)).1 += 1;
        }
        map.values()
            .into_iter()
            .fold(0, |ans, pair| ans + pair.0 * pair.1)
    }
}

fn main() {}
