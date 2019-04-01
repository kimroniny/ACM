class Solution:
    def maxSubArray(self, nums: 'List[int]') -> 'int':
        result = nums[0]
        s = 0
        for x in nums:
            s += x
            result = max(result, s)
            if s <= 0:
                s = 0
                
        return result
if __name__ == "__main__":
    print(
        Solution().maxSubArray(
            [-1, 1, -1]
        )
    )    
