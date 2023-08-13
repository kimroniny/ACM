
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val: val
    }
  }
}
struct Solution{}
impl Solution {
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        use std::collections::BinaryHeap;
        use std::cmp::Reverse;
        let mut heap = BinaryHeap::new();
        for row in lists {
            let mut point = row;
            while let Some(node) = point {
                heap.push(Reverse(node.val));
                point = node.next;
            }            
        }

        let mut head:Option<Box<ListNode>> = None;
        let mut node = &mut head;
        while let Some(val) = heap.pop() {
          let temp: Box<ListNode> = Box::new(ListNode::new(val.0));
          *node = Some(temp);
          node = &mut (*node).as_mut().unwrap().next;
        }
        head
    }
}
fn main() {
    let lists = vec![
      Some(Box::new(ListNode{val:1, next:Some(Box::new(ListNode{val:2, next:Some(Box::new(ListNode{val:3, next:None}))}))})),
      Some(Box::new(ListNode{val:4, next:Some(Box::new(ListNode{val:5, next:Some(Box::new(ListNode{val:6, next:None}))}))})),
      Some(Box::new(ListNode{val:7, next:Some(Box::new(ListNode{val:8, next:Some(Box::new(ListNode{val:9, next:None}))}))})),
    ];
    let ans = Solution::merge_k_lists(lists);
    println!("{:#?}", ans);
    println!("Hello, world!");
}
