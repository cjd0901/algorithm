struct Solution {}

// 1653. 使字符串平衡的最少删除次数
impl Solution {
    pub fn minimum_deletions(s: String) -> i32 {
        let (mut a, mut b, len) = (0, 0, s.len());
        for ch in s.chars().into_iter() {
            match ch {
                'a' => a += 1,
                'b' => b = std::cmp::max(a, b) + 1,
                _ => {}
            }
        }
        len as i32 - std::cmp::max(a, b)
    }
}

fn main() {}
