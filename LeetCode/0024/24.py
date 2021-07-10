# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def swapPairs(self, head: 'ListNode') -> 'ListNode':
        head2 = head
        if head.next is None:
            return head
        result_head = ListNode(head.next.val)
        result_head.next = head
        last = None
        while head2 is not None:
            second = head2.next
            if second is None:
                break
            replace = ListNode(second.val)
            replace.next = head2
            head2.next = second.next
            second = replace
            if last is not None:
                last.next = second
            last = second.next
            head2 = head2.next
        return result_head

if __name__ == "__main__":
    arr = [
        1, 2, 3, 4
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

    head = Solution().swapPairs(
        headx
    )
    while head is not None:
        print(head.val)
        head = head.next