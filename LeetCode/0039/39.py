class Solution:
    result = []
    result_all = []
    def combinationSum(self, candidates: 'List[int]', target: 'int') -> 'List[List[int]]':
        self.result = []
        self.result_all = []
        self.do(target, candidates)
        return self.result_all
    
    def do(self, p, nums):
        for i in range(len(nums)):
            if p < nums[i]:
                return 
            elif p == nums[i]:
                self.result_all.append([x for x in self.result]+[p])
                return
            else:
                self.result.append(nums[i])
                self.do(p-nums[i], nums[i:])
                self.result.pop()

if __name__ == "__main__":
    print(
        Solution().combinationSum(
            [2, 3, 6, 7], 7
        )
    )
