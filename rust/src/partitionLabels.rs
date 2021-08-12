// Partition Labels
// leetcode: https://leetcode-cn.com/problems/partition-labels/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let s = s.as_bytes();
        let mut map:HashMap<u8, i32> = HashMap::new();
        let mut ans = vec![];
        for (i, b) in s.iter().enumerate() {
            *map.entry(*b).or_insert(i as i32) = i as i32;
        }
        let (mut start, mut end) = (0 as usize, 0 as usize);
        for (i, b) in s.iter().enumerate() {
            if let Some(v) = map.get(b) {
                if *v > end as i32 {
                    end = *v as usize;
                }
            }
            if i == end {
                ans.push((end-start+1) as i32);
                start = end+1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::partition_labels(String::from("ababcbacadefegdehijhklij"));
    println!("{:?}", ans);
}