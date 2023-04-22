struct Solution {}

// 1039. 多边形三角剖分的最低得分
impl Solution {
    pub fn min_score_triangulation(values: Vec<i32>) -> i32 {
        let len = values.len();
        let mut cache = vec![vec![-1; len]; len];
        fn dfs(values: &Vec<i32>, cache: &mut Vec<Vec<i32>>, i: usize, j: usize) -> i32 {
            if i + 1 == j {
                return 0;
            }
            if cache[i][j] != -1 {
                return cache[i][j];
            }
            let mut res = i32::MAX;
            for k in i + 1..j {
                res = res.min(
                    dfs(values, cache, i, k)
                        + dfs(values, cache, k, j)
                        + values[i] * values[j] * values[k],
                );
            }
            cache[i][j] = res;
            res
        }
        dfs(&values, &mut cache, 0, len - 1)
    }
}

fn main() {}
