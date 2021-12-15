// Loud and Rich
// leetcode: https://leetcode-cn.com/problems/loud-and-rich/
struct Solution {}

impl Solution {
    pub fn loud_and_rich(richer: Vec<Vec<i32>>, quiet: Vec<i32>) -> Vec<i32> {
        let len = quiet.len();
        let mut edges: Vec<Vec<i32>> = vec![vec![]; len];
        let mut ans = vec![-1; len];
        for rich in richer.iter() {
            edges[rich[1] as usize].push(rich[0]);
        }
        for i in 0..len {
            search_min(&edges, &mut ans, &quiet, i);
        }
        ans
    }
}

fn search_min(edges: &Vec<Vec<i32>>, ans: &mut Vec<i32>, quiet: &Vec<i32>, idx: usize) {
    if ans[idx] != -1 {
        return
    }
    ans[idx] = idx as i32;
    for edge in &edges[idx] {
        search_min(&edges, ans, &quiet, idx);
        if quiet[ans[*edge as usize] as usize] < quiet[ans[idx] as usize] {
            ans[idx] = ans[*edge as usize];
        }
    }
}

fn main() {
    let mut v = Vec::new();
    v.push(vec![0,2]);
    v.push(vec![1,2]);
    let ans = Solution::loud_and_rich(v, vec![0,1,2]);
    println!("{:?}", ans);
}