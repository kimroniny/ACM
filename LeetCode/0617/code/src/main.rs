// Definition for a binary tree node.
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
use std::borrow::BorrowMut;
use std::ops::Deref;
use std::rc::{Rc};
use std::cell::RefCell;
struct Solution {}
impl Solution {
    pub fn merge_trees(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        
        if root1 == None && root2 == None {
            return None;
        }

        let mut val = 0;
        let mut left1:Option<Rc<RefCell<TreeNode>>> = None;
        let mut right1:Option<Rc<RefCell<TreeNode>>> = None;
        let mut left2:Option<Rc<RefCell<TreeNode>>> = None;
        let mut right2:Option<Rc<RefCell<TreeNode>>> = None;
        if let Some(node1) = root1.clone() {
            val += node1.borrow().val;
            left1 = (*node1).borrow_mut().left.take();
            right1 = (*node1).borrow_mut().right.take();
        }
        if let Some(node2) = root2.clone() {
            val += node2.borrow().val;
            left2 = (*node2).borrow_mut().left.take();
            right2 = (*node2).borrow_mut().right.take();
        }
        
        Some(Rc::new(RefCell::new(TreeNode {
          val: val,
          left: Solution::merge_trees(left1, left2),
          right: Solution::merge_trees(right1, right2)
        })))
    }
}
fn main() {
    println!("Hello, world!");
}
