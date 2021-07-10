# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        prehead = ListNode(0)
        prehead.next = head
        cnt = 1
        p, prep = head, prehead
        while cnt < m:
            p, prep, cnt = p.next, prep.next, cnt+1
        np, p = p, p.next
        while cnt < n:
            pnext = p.next
            p.next = np
            np = p
            p = pnext
            cnt += 1
        prep.next.next = p
        prep.next = np
                    
        return prehead.next
        

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
        headx, 1, 5
    )
    while head is not None:
        print(head.val)
        head = head.next

            