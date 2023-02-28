struct Solution {}

use std::collections::HashMap;

//2363.合并相似的物品
impl Solution {
    pub fn merge_similar_items(items1: Vec<Vec<i32>>, items2: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut ans = items1
            .into_iter()
            .chain(items2)
            .fold(HashMap::new(), |mut m, item| {
                let e = m.entry(item[0]).or_insert(0);
                *e += item[1];
                m
            })
            .into_iter()
            .map(|(value, weight)| vec![value, weight])
            .collect::<Vec<Vec<i32>>>();
        ans.sort_by(|a, b| a[0].cmp(&b[0]));
        ans
    }
}

fn main() {}
