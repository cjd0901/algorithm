// Find the Town Judge
// leetcode: https://leetcode-cn.com/problems/find-the-town-judge/
struct Solution {}

impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut out_degrees = vec![0; (n+1) as usize];
        let mut in_degrees = vec![0; (n+1) as usize];
        for tr in trust.iter() {
            in_degrees[tr[0] as usize] += 1;
            out_degrees[tr[1] as usize] += 1;
        }
        let mut q = Vec::new();
        for i in 1..(n+1) as usize {
            if in_degrees[i] == 0 {
                q.push(i as i32);
            }
        }
        for i in q.iter() {
            if out_degrees[*i as usize] == n-1 {
                return *i;
            }
        }
        -1
    }
}

fn main() {
    let mut v = Vec::new();
    v.push(vec![1,2]);
    let ans = Solution::find_judge(2, v);
    println!("{:?}", ans);
}