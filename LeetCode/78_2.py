# binary method
class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        candidates=[]
        res=[]
        for i in range(2**len(nums)):
            candidates.append(bin(i))
        for b in candidates:
            tmp=[]
            b=b[2:]
            b=list(b[::-1])
            for i,k in enumerate(b):
                if k=='1':
                    tmp.append(nums[i])
            res.append(tmp)
        return res