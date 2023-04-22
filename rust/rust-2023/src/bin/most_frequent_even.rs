struct Solution {}

// 2404. 出现最频繁的偶数元素
impl Solution {
    pub fn most_frequent_even(nums: Vec<i32>) -> i32 {
        let mut m = std::collections::HashMap::new();
        for num in nums.into_iter() {
            if num % 2 == 0 {
                *m.entry(num).or_insert(0) += 1;
            }
        }
        m.into_iter()
            .fold((-1, 0), |ans, x| {
                if x.1 > ans.1 {
                    (x.0, x.1)
                } else if x.1 == ans.1 {
                    (x.0.min(ans.0), x.1)
                } else {
                    ans
                }
            })
            .0
    }
}

fn main() {}
