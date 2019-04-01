# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseBetween(self, head: 'ListNode', m: 'int', n: 'int') -> 'ListNode':
        if m == n:
            return head
        last = ListNode(0)
        last.next = head
        last0 = last
        cnt = 1
        head0 = head
        while cnt < m:
            cnt += 1
            last0 = head0
            head0 = head0.next
        headm = head0
        while cnt < n:
            cnt += 1
            head0 = head0.next
        headn = head0
        nextdat = headn.next
        last0.next = headn
        second = headm.next
        headm.next = nextdat
        # return last.next

        last2 = headm
        while second != headn:
            next0 = second.next
            second.next = last2
            last2 = second
            second = next0
        
        return last.next

if __name__ == "__main__":
    arr = [
        1, 2, 3, 4, 5
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

    head = Solution().reverseBetween(
        headx, 2, 3
    )
    while head is not None:
        print(head.val)
        head = head.next

            