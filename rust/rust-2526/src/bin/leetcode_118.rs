struct Solution {}

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut ans = vec![vec![]; num_rows as usize];
        for i in 0..num_rows as usize {
            ans[i] = vec![1; i+1];
            if i > 1 {
                for j in 1..i {
                    ans[i][j] = ans[i-1][j] + ans[i-1][j-1];
                }
            }
        }
        ans
    }
}

fn main() {

}