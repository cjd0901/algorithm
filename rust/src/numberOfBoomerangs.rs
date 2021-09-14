// Number of Boomerangs
// leetcode: https://leetcode-cn.com/problems/number-of-boomerangs/
use std::collections::HashMap;

struct Solution {}
impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        for (i, p) in points.iter().enumerate() {
            let mut map: HashMap<i32, i32> = HashMap::new();
            for (j, q) in points.iter().enumerate() {
                if i == j {
                    continue;
                }
                let dis = (q[0]-p[0]) * (q[0]-p[0]) + (q[1]-p[1]) * (q[1]-p[1]);
                *map.entry(dis).or_insert(0) += 1;
            }
            for n in map.values() {
                ans += n * (n-1);
            }
        }
        ans
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![0, 0]);
    v.push(vec![1, 0]);
    v.push(vec![2, 0]);
    let ans = Solution::number_of_boomerangs(v);
    println!("{}", ans);
}