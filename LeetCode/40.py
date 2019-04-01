class Solution:
    result = []
    result_all = []
    def combinationSum2(self, candidates: 'List[int]', target: 'int') -> 'List[List[int]]':
        self.result = []
        self.result_all = []
        candidates = sorted(candidates)
        self.do(target, candidates)
        return self.result_all
    
    def do(self, p, nums):
        for i in range(len(nums)):
            if p < nums[i]:
                return 
            elif p == nums[i]:
                s = [x for x in self.result]+[p]
                if s not in self.result_all:
                    self.result_all.append(s)
                return
            else:
                self.result.append(nums[i])
                self.do(p-nums[i], nums[i+1:])
                self.result.pop()

if __name__ == "__main__":
    print(
        Solution().combinationSum2(
            [10,1,2,7,6,1,5], 8
        )
    )
