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
    pub fn evaluate_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let r = root.unwrap();
        let left = r.borrow_mut().left.take();
        let right = r.borrow_mut().right.take();
        let res = match r.borrow().val {
            0 => false,
            1 => true,
            2 => Self::evaluate_tree(left) || Self::evaluate_tree(right),
            _ => Self::evaluate_tree(left) && Self::evaluate_tree(right),
        };
        res
    }
}

fn main() {}
