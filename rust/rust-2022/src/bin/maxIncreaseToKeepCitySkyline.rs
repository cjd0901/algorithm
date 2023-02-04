// Max Increase to Keep City Skyline
// leetcode: https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/
struct Solution {}

impl Solution {
    pub fn max_increase_keeping_skyline(grid: Vec<Vec<i32>>) -> i32 {
        let len = grid.len();
        let (mut vertical, mut horizontal) = (vec![0; len], vec![0; len]);
        let mut ans = 0;
        for i in 0..len {
            for j in 0..len {
                let skyline = grid[i][j];
                if skyline > horizontal[i] {
                    horizontal[i] = skyline;
                }
                if skyline > vertical[j] {
                    vertical[j] = skyline;
                }
            }
        }
        for i in 0..len {
            for j in 0..len {
                let skyline = grid[i][j];
                ans += std::cmp::min(vertical[j] - skyline, horizontal[i] - skyline);
            }
        }
        ans
    }
}

fn main() {
    let mut v = Vec::new();
    v.push(vec![3,0,4,8]);
    v.push(vec![2,4,5,7]);
    v.push(vec![9,2,6,3]);
    v.push(vec![0,3,1,0]);
    let ans = Solution::max_increase_keeping_skyline(v);
    println!("{}", ans);
}