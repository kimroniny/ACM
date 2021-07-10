class Solution:
    def canJump(self,nums:'list[int]')->'bool':
        l=len(nums)
        idx=l-1
        for i in range(l-2,-1,-1):
            if idx-i<=nums[i]:
                idx=i
        return idx==0