# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists: 'List[ListNode]') -> 'ListNode':
        import heapq
        result = []
        for node in lists:
            while node is not None:
                heapq.heappush(result, node.val)
                node = node.next
        head = ListNode(-1)
        p = head
        for i in range(len(result)):
            newNode = ListNode(heapq.heappop(result))
            p.next = newNode
            p = newNode
        return head.next

# if __name__ == "__main__":
#     head = Solution(
