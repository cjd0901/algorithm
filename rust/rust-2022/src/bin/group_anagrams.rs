// Group Anagrams
// leetcode: https://leetcode-cn.com/problems/group-anagrams-lcci/
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut ans = vec![];
        let mut map: HashMap<String, Vec<String>> = HashMap::new();
        for s in strs.iter() {
            let mut temp = s.clone().into_bytes();
            temp.sort();
            let b = String::from_utf8(temp).unwrap();
            map.entry(b).or_insert(vec![]).push(s.to_string());
        }
        for v in map.values() {
            ans.push(v.clone());
        }
        ans
    }
}

fn main() {
    let mut v = vec![];
    v.push(String::from("eat"));
    v.push(String::from("tea"));
    v.push(String::from("tan"));
    v.push(String::from("ate"));
    v.push(String::from("nat"));
    v.push(String::from("bat"));
    let ans = Solution::group_anagrams(v);
    println!("{:?}", ans);
}