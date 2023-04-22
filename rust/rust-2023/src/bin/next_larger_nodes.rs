struct Solution {}

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

// 1019. 链表中的下一个更大节点
impl Solution {
    pub fn next_larger_nodes(head: Option<Box<ListNode>>) -> Vec<i32> {
        fn dfs(ln: Option<Box<ListNode>>, i: usize, ans: &mut Vec<i32>, stack: &mut Vec<i32>) {
            if ln.is_none() {
                return;
            }
            ans.push(0);
            let ln = ln.unwrap();
            dfs(ln.next, i + 1, ans, stack);
            while stack.len() > 0 && *stack.last().unwrap() <= ln.val {
                stack.pop();
            }
            if stack.len() > 0 {
                ans[i] = *stack.last().unwrap();
            }
            stack.push(ln.val);
        }
        let mut stack = vec![];
        let mut ans = vec![];
        dfs(head, 0, &mut ans, &mut stack);
        ans
    }
}

fn main() {}
