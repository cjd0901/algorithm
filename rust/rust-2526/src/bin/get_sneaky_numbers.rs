use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn get_sneak_numbers(nums: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![];
        let mut set: HashSet<i32> = HashSet::new();
        for num in nums {
            if set.get(&num).is_some() {
                ans.push(num);
            }
            set.insert(num);
        }
        ans
    }
}

fn main() {

}