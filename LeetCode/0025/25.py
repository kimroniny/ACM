# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseKGroup(self, head: 'ListNode', k: 'int') -> 'ListNode':
        last = ListNode(0)
        newhead = last
        node = head
        while node is not None:
            (last, nextbat, reverse) = self.dok(node, k, last, k)
            node = nextbat
            if nextbat is None and reverse:
                last.next = None
        return newhead.next
    
    def dok(self, node, k, last, maxk):
        if k==0:
            return (None, node, True)
        if node is None:
            return (None, None, False)
        (p, nextbat, reverse) = self.dok(node.next, k-1, last, maxk)
        if k == 1 or (not reverse and k == maxk):
            last.next = node
        if k == maxk:
            last = node
        if k >= 2 and reverse:
            node.next.next = node
        return (last, nextbat, reverse)

if __name__ == "__main__":
    arr = [
        1, 2
    ]
    last = None
    headx = None
    for x in arr:
        head2 = ListNode(x)
        if last is None:
            headx = head2
        else:
            last.next = head2
        last = head2

    head = Solution().reverseKGroup(
        headx, 2
    )
    while head is not None:
        print(head.val)
        head = head.next