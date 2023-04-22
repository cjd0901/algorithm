struct Solution {}

// 1023. 驼峰式匹配
impl Solution {
    pub fn camel_match(queries: Vec<String>, pattern: String) -> Vec<bool> {
        let len = queries.len();
        queries.into_iter().fold(vec![], |mut ans, query| {
            let mut y = 0;
            let pattern = pattern.chars().collect::<Vec<_>>();
            for ch in query.chars().into_iter() {
                if y < pattern.len() && ch == pattern[y] {
                    y += 1;
                } else if ch.is_uppercase() {
                    ans.push(false);
                    return ans;
                }
            }
            ans.push(if y < pattern.len() { false } else { true });
            ans
        })
    }
}

fn main() {}
