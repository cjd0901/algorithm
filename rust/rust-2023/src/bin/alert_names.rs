struct Soltuion {}

use std::collections::HashMap;
use std::ops::Deref;
use std::str::FromStr;

// 1604.警告一小时内使用相同员工卡大于等于三次的人
impl Solution {
    pub fn alert_names(key_name: Vec<String>, key_time: Vec<String>) -> Vec<String> {
        let mut ans = Vec::new();
        let mut key_map = HashMap::new();
        for (i, name) in key_name.iter().enumerate() {
            let key = key_map.entry(name).or_insert(Vec::new());
            key.push(key_time.get(i).unwrap());
        }
        for (name, times) in key_map.iter() {
            let mut times: Vec<u32> = times
                .iter()
                .map(|time| {
                    let time: Vec<&str> = time.split(":").collect();
                    let hours: u32 = FromStr::from_str(time[0]).unwrap();
                    let minutes: u32 = FromStr::from_str(time[1]).unwrap();
                    hours * 60 + minutes
                })
                .collect();
            times.sort();
            if times.len() < 3 {
                continue;
            }
            let mut j = 0;
            while j + 2 < times.len() {
                if times[j + 2] - times[j] <= 60 {
                    ans.push(name.deref().clone());
                    break;
                }
                j += 1;
            }
        }
        ans.sort();
        ans
    }
}

fn main() {}
