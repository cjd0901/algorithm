// 传递信息
// leetcode: https://leetcode-cn.com/problems/chuan-di-xin-xi/
struct Solution {}
impl Solution {
    pub fn num_ways(n: i32, relation: Vec<Vec<i32>>, k: i32) -> i32 {
        let mut rel: Vec<Vec<i32>> = vec![vec![]; n as usize];
        let mut count = 0;
        for i in 0..relation.len() {
            let r = &relation[i];
            rel[r[0] as usize].push(r[1]);
        }
        transfer(0, &rel, &mut count, 0, k);
        count
    }
}

fn transfer(fre: i32, rel: &Vec<Vec<i32>>, count: &mut i32, cur: i32, k: i32) {
    if fre == k {
        if cur == rel.len() as i32 - 1 {
            *count += 1;
        }
        return
    }else {
        for n in &rel[cur as usize] {
            transfer(fre + 1, rel, count, *n, k)
        }
    }
}

fn main() {
    let mut relation = vec![];
    relation.push(vec![0, 2]);
    relation.push(vec![2, 1]);
    relation.push(vec![3, 4]);
    relation.push(vec![2, 3]);
    relation.push(vec![1, 4]);
    relation.push(vec![2, 0]);
    relation.push(vec![0, 4]);
    let count = Solution::num_ways(5, relation, 3);
    println!("{}", count);
}