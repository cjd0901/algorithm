struct Solution {}

impl Solution {
    pub fn permutation(mut s: String) -> Vec<String> {
        let l = s.len();
        let mut res = vec![];
        let mut vis = vec![false; l];
        let mut t = String::with_capacity(l);
        let a = unsafe{ s.as_bytes_mut() };
        a.sort();
        fn backtrace(pos: usize, res: &mut Vec<String>, t: &mut String, vis: &mut Vec<bool>, l: usize, a: &[u8]) {
            if pos == l {
                res.push(t.clone());
                return;
            }
            for i in 0..l {
                if vis[i] || i > 0 && !vis[i-1] && a[i-1] == a[i] {
                    continue
                }
                t.push(a[i] as char);
                vis[i] = true;
                backtrace(pos + 1, res, t, vis, l, a);
                t.pop();
                vis[i] = false;
            }
        }
        backtrace(0, &mut res, &mut t, &mut vis, l, &a);
        res
    }
}

fn main() {
    let ans = Solution::permutation(String::from("abc"));
    println!("{:?}", ans);
}