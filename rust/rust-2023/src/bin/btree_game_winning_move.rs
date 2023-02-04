use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

struct Solution {}

impl Solution {
    pub fn btree_game_winning_move(root: Option<Rc<RefCell<TreeNode>>>, n: i32, x: i32) -> bool {
        let mut ans = 0;
        fn dfs(node: Option<Rc<RefCell<TreeNode>>>, n: i32, x: i32, ans: &mut i32) -> i32 {
            match node {
                Some(r) => {
                    let ls = dfs(r.borrow_mut().left.take(), n, x, ans);
                    let rs = dfs(r.borrow_mut().right.take(), n, x, ans);
                    if r.borrow().val == x {
                        *ans = *[ls, rs, n - ls - rs - 1].iter().max().unwrap()
                    }
                    1 + ls + rs
                }
                None => 0,
            }
        }
        dfs(root, n, x, &mut ans);
        ans > n - ans
    }
}

fn main() {}
