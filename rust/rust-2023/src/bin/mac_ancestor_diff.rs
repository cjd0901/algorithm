use std::cell::RefCell;
use std::rc::Rc;

struct Solution {}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

impl Solution {
    pub fn max_ancestor_diff(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        fn dfs(tn: Option<Rc<RefCell<TreeNode>>>, ans: &mut i32, ma: i32, mi: i32) {
            if tn.is_none() {
                return;
            }
            let tn = tn.unwrap();
            let val = tn.borrow().val;
            let ma = std::cmp::max(ma, val);
            let mi = std::cmp::min(mi, val);
            *ans = std::cmp::max(*ans, std::cmp::max((ma - val).abs(), (mi - val).abs()));
            dfs(tn.borrow_mut().left.take(), ans, ma, mi);
            dfs(tn.borrow_mut().right.take(), ans, ma, mi);
        }
        let mut ans = 0;
        dfs(root, &mut ans, i32::MIN, i32::MAX);
        ans
    }
}

fn main() {}
