class Solution:
    nums = []
    result = []
    result_all = []
    def subsets(self, nums: 'List[int]') -> 'List[List[int]]':
        self.nums = nums
        self.result, self.result_all = [], []
        self.do(0)
        return self.result_all + ([[]] if len(self.result_all) > 0 else [])

    def do(self, r):
        while r < len(self.nums):
            self.result.append(self.nums[r])
            self.result_all.append([x for x in self.result])
            r += 1
            self.do(r)
            self.result.pop()

if __name__ == "__main__":
    print(
        Solution().subsets(
            [1, 2, 3]
        )
    )