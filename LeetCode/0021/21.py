# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1: 'ListNode', l2: 'ListNode') -> 'ListNode':
        result = None
        head = ListNode(0)
        last = head
        while l1 is not None or l2 is not None:
            if l2 is None or l1.val < l2.val:
                result = ListNode(l1.val)
                l1 = l1.next
            else:
                result = ListNode(l2.val)
                l2 = l2.next
            last.next = result
            last = result
        return head.next

if __name__ == "__main__":
    pass