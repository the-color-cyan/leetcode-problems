// Definition for singly-linked list.
class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val?: number, next?: ListNode | null) {
        this.val = val === undefined ? 0 : val;
        this.next = next === undefined ? null : next;
    }
}

function addTwoNumbers(
    l1: ListNode | null,
    l2: ListNode | null,
): ListNode | null {
    const dummyHead = new ListNode();
    let node: ListNode = dummyHead;
    let carry = 0;

    while (l1 || l2 || carry) {
        const val_1 = l1 ? l1.val : 0;
        const val_2 = l2 ? l2.val : 0;
        const sum = val_1 + val_2 + carry;

        carry = Math.floor(sum / 10);
        node.next = new ListNode(sum % 10);
        node = node.next;

        l1 = l1 ? l1.next : l1;
        l2 = l2 ? l2.next : l2;
    }

    return dummyHead.next;
}
