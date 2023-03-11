struct Solution {}

// 2379. 得到 K 个黑块的最少涂色次数
impl Solution {
    pub fn minimum_recolors(blocks: String, k: i32) -> i32 {
        let blocks: Vec<char> = blocks.chars().collect();
        let (mut start, mut end) = (0, 0);
        let mut ops = 0;
        for i in 0..k as usize {
            if blocks[i] == 'W' {
                ops += 1;
            }
            end += 1;
        }
        let mut ans = ops;
        while end < blocks.len() {
            if blocks[end] == 'W' {
                ops += 1;
            }
            if blocks[start] == 'W' {
                ops -= 1;
            }
            start += 1;
            end += 1;
            ans = ans.min(ops);
        }
        ans
    }
}

fn main() {}
