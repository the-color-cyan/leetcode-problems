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

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut l1, mut l2) = (l1.as_ref(), l2.as_ref());

        let mut dummyHead = Box::new(ListNode::new(0));
        let mut node = dummyHead.as_mut();
        let mut carry: i32 = 0;

        while (l1.is_some() || l2.is_some() || carry != 0) {
            let val_1: i32 = if (l1.is_some()) { l1.unwrap().val } else { 0 };
            let val_2: i32 = if (l2.is_some()) { l2.unwrap().val } else { 0 };

            let sum: i32 = val_1 + val_2 + carry;

            carry = sum / 10;
        }

        return dummyHead.next;
    }
}
