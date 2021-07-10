# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head is None: return head
        mid = head
        pre = ListNode(mid.val-1)
        prehead = pre
        pre.next = mid
        back = mid.next
        while back is not None:
            while back is not None and back.val == mid.val:
                back = back.next
            if mid.next == back:
                pre = mid
            mid = back
            if back != None: back = back.next
            pre.next = mid
        return prehead.next


if __name__ == "__main__":
    values = [
        1, 1, 1, 4, 5
    ]
    head = ListNode(values[0])
    p = head
    for val in values[1:]:
        newp = ListNode(val)
        p.next = newp
        p = newp
    
    res = Solution().deleteDuplicates(head)
    while res is not None:
        print(res.val,sep=" ")
        res = res.next
