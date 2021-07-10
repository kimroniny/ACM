# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if root is None: return True
        return self.do(root.left, root.right)
    
    def do(self, left, right):
        if (left is None and right is None): return True
        if left is None or right is None or left.val != right.val: return False
        return (
            self.do(left.left, right.right) and 
            self.do(left.right, right.left)
        )
    
    def test(self) -> bool:
        pass

def buildTree(treelist: list, node: TreeNode, num: int):
    if num *2 <= len(treelist):
        node.left = TreeNode(treelist[num*2-1])
        buildTree(treelist, node.left, num*2)
    if num *2 +1 <= len(treelist):
        node.right = TreeNode(treelist[num*2])
        buildTree(treelist, node.right, num*2+1)

def printTree(node: TreeNode):
    if node is None: return
    print(node.val, end="\n")
    printTree(node.left)
    printTree(node.right)

if __name__ == "__main__":
    treelist = [1, 2, 2, 3, 4, 4, 3]
    root = TreeNode(treelist[0])
    buildTree(treelist, root, 1)
    # printTree(root)
    print(
        Solution().isSymmetric(
            root
        )
    )
    # print(Solution().test())


