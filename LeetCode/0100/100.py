# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSameTree(self, p: 'TreeNode', q: 'TreeNode') -> 'bool':
        if p is None and q is None: return True
        if p is not None and q is not None and p.val == q.val:
            return self.isSameTree(p.left, q.left) & self.isSameTree(p.right, q.right)
        else:
            return False

if __name__ == "__main__":
    
    
