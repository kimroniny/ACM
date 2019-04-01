class Solution:
    result = None
    nums = None
    result_all = None
    def subsetsWithDup(self, nums: 'List[int]') -> 'List[List[int]]':
        self.result = []
        self.result_all = []
        self.nums = nums
        if len(nums) == 0: return []
        self.do(0)
        return self.result_all
    
    def do(self, x):
        if x > 0 and x < len(self.nums) and self.nums[x-1] == self.nums[x]: return 
        while x < len(self.nums):
            self.result.append(self.nums[x])
            self.result_all.append([x for x in self.result])
            x += 1
            self.do(x)
        self.result.pop()

if __name__ == "__main__":
    print(
        Solution().subsetsWithDup(
            [1, 2, 2]
        )
    )

        