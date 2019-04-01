# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head is None or (head.next is None) or (head.next.val != head.val and head.next.next is None):
            return head
        if head.next.val == head.val and head.next.next is None: return None
        node0, node1 = head, head.next
        while node1.next is not None:
            flag = True
            while node1.next is not None and node1.next.val == node1.val:
                flag = False
                node0.next = node1.next
                node1 = node1.next
            if node1.next is None:
                node0.next = None
            else:
                if not flag: 
                    node0.next = node1.next
                    node1 = node1.next
                else:

            

        