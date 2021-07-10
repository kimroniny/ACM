class Solution:
    v = None
    result = None
    result_all = None
    def permuteUnique(self, nums: 'List[int]') -> 'List[List[int]]':
        lens = len(nums)
        self.v = [False for i in range(lens)]
        self.result = []
        self.result_all = []
        self.do(nums)
        return self.result_all

    def do(self, nums):
        flag = True
        for i in range(len(nums)):
            if not self.v[i]:
                flag = False
                self.v[i] = True
                self.result.append(nums[i])
                # print(nums[i])
                self.do(nums)
                self.v[i] = False
                self.result.pop()
        if flag:
            s = [x for x in self.result]
            if s not in self.result_all:
                self.result_all.append(s)