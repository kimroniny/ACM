class Solution:
    v = None
    result = None
    result_all = None
    def permute(self, nums: 'List[int]') -> 'List[List[int]]':
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
            self.result_all.append([x for x in self.result])

if __name__ == "__main__":
    print(
        Solution().permute(
            [1, 2, 3]
        )
    )

            