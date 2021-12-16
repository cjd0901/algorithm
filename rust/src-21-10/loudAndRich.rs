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

    // 拓扑排序
    pub fn loud_and_rich2(richer: Vec<Vec<i32>>, quiet: Vec<i32>) -> Vec<i32> {
        let len = quiet.len();
        let mut edges: Vec<Vec<i32>> = vec![vec![]; len];
        let mut in_degrees = vec![0; len];
        for rich in richer.iter() {
            edges[rich[0] as usize].push(rich[1]);
            in_degrees[rich[1] as usize] += 1;
        }
        let mut ans = vec![0; len];
        for i in 0..len {
            ans[i] = i as i32;
        }
        let mut q = Vec::new();
        for (i, in_degree) in in_degrees.iter().enumerate() {
            if *in_degree == 0 {
                q.push(i as i32);
            }
        }
        while q.len() > 0 {
            let t = q.pop().unwrap() as usize;
            for edge in edges[t].iter() {
                let edge = *edge;
                if quiet[ans[t] as usize] < quiet[ans[edge as usize] as usize] {
                    ans[edge as usize] = ans[t];
                }
                in_degrees[edge as usize] -= 1;
                if in_degrees[edge as usize] == 0 {
                    q.push(edge);
                }
            }
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
    let ans = Solution::loud_and_rich2(v, vec![0,1,2]);
    println!("{:?}", ans);
}