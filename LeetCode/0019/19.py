# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: 'ListNode', n: 'int') -> 'ListNode':
        cnt = 0
        head2 = head
        while head2 is not None:
            cnt += 1
            head2 = head2.next
        head2 = head
        cnt2 = 0
        if cnt2 == cnt-n and n != 0:
            head = head.next
        else:
            while head2 is not None:
                cnt2 += 1
                if cnt2 == cnt-n and n != 0:
                    head2.next = head2.next.next
                    break
                head2 = head2.next
        return head

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

    head = Solution().removeNthFromEnd(
        headx, 2
    )
    while head is not None:
        print(head.val)
        head = head.next
