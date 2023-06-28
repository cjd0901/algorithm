struct Solution {}

impl Solution {
    pub fn can_make_pali_queries(s: String, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let s_bytes = s.as_bytes();
        let mut sum = vec![0; s_bytes.len() + 1];
        for (i, b) in s_bytes.into_iter().enumerate() {
            sum[i + 1] = sum[i] ^ (1 << (*b - b'a'));
        }
        queries
            .into_iter()
            .map(|query| {
                let mut cnt = 0;
                let (l, r, k) = (query[0] as usize, query[1] as usize, query[2]);
                let sum_xor = sum[r + 1] ^ sum[l];
                for i in 0..26 {
                    cnt += sum_xor >> i & 1
                }
                cnt / 2 <= k
            })
            .collect()
    }
}

fn main() {}
