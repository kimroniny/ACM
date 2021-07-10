class Solution:
    result_all = None
    def subsetsWithDup(self, nums: 'List[int]') -> 'List[List[int]]':
        self.result_all = []
        if not nums: return []
        nums.sort()
        self.do(nums, [])
        return self.result_all
    
    def do(self, nums, result):
        self.result_all.append(result)
        for idx, val in enumerate(nums):
            if idx > 0 and nums[idx] == nums[idx-1]: continue
            self.do(nums[idx+1:], result + [val])

if __name__ == "__main__":
    print(
        Solution().subsetsWithDup(
            [1, 2, 2]
        )
    )

        