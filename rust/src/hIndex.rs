// H Index
// leetcode: https://leetcode-cn.com/problems/h-index/
struct Solution {}
impl Solution {
    pub fn h_index(citations: Vec<i32>) -> i32 {
        let (mut l, mut r) = (0 as usize, citations.len());
        while l < r {
            let mid = l + (r - l + 1) / 2;
            if check(&citations, mid as i32) { l = mid }
            else { r = mid - 1 }
        }
        r as i32
    }
    pub fn h_index2(citations: Vec<i32>) -> i32 {
        let mut citations = citations;
        citations.sort();
        let (mut h, mut i) = (0, (citations.len() - 1) as i32);
        while i >= 0 && citations[i as usize] > h {
            h += 1;
            i -= 1;
        }
        h
    }
}

fn check(citations: &Vec<i32>, mid: i32) -> bool {
    let mut ans = 0;
    for c in citations.iter() {
        if *c >= mid {
            ans += 1;
        }
    }
    ans >= mid
}

fn main() {
    let h = Solution::h_index2(vec![0]);
    println!("{}", h);
}