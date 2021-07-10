# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: 'ListNode', n: 'int') -> 'ListNode':
        head2 = head
        head3 = head
        cnt = 0
        if n == 0:
            return head
        while head2 is not None:
            cnt += 1
            if cnt > n+1:
                head3 = head3.next
            head2 = head2.next
        if n == cnt:
            head = head.next
        else:
            head3.next = head3.next.next
        return head




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

    head = Solution().removeNthFromEnd(
        headx, 2
    )
    while head is not None:
        print(head.val)
        head = head.next
