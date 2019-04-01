# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def rotateRight(self, head: 'ListNode', k: 'int') -> 'ListNode':
        headp = head
        cnt = 0
        last = None
        while headp is not None:
            cnt += 1
            if headp.next is None:
                last = headp
            headp = headp.next
        
        k %= cnt
        if k == 0: return head

        last.next = head
        headp = head
        cnt0 = 0
        while cnt0 < cnt-k - 1:
            headp = headp.next
            cnt0 += 1
        head = headp.next
        headp.next = None
        return head

if __name__ == "__main__":

    a = [1, 2, 3, 4, 5, 6]
    head = ListNode(-1)
    head0 = head
    for x in a:
        cur = ListNode(x)
        head0.next = cur
        head0 = cur
    
    result = Solution().rotateRight(
        head.next, 6
    )

    while result is not None:
        print(result.val)
        result = result.next

    

        
        
        
        