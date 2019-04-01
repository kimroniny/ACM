# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteDuplicates(self, head: 'ListNode') -> 'ListNode':
        if head is None: return None
        
        head0 = ListNode(head.val-1)
        head0.next = head
        last = head0
        headp = head

        while headp is not None:
            if headp.val == last.val:
                last.next = headp.next
            else:
                last = headp
            headp = headp.next
        
        return head0.next

if __name__ == "__main__":

    a = [1, 1]
    head = ListNode(-1)
    head0 = head
    for x in a:
        cur = ListNode(x)
        head0.next = cur
        head0 = cur
    
    result = Solution().deleteDuplicates(
        head.next
    )

    while result is not None:
        print(result.val)
        result = result.next
