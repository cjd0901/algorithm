use std::rc::Rc;
use std::cell::RefCell;
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
      right: None
    }
  }
}

struct Solution {}

impl Solution {
    pub fn bst_to_gst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        pub fn dfs(tn: Option<Rc<RefCell<TreeNode>>>, sum: &mut i32) {
            if tn.is_none() {
                return;
            }
            let tn = tn.unwrap();
            dfs(tn.borrow().right.clone(), sum);
            *sum += tn.borrow().val;
            tn.borrow_mut().val = *sum;
            dfs(tn.borrow().left.clone(), sum);
        }
        let mut sum = 0;
        dfs(root.clone(), &mut sum);
        root
    }
}

fn main() {}