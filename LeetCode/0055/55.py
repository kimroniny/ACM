class Solution:
    f = None
    nums = None
    def canJump(self, nums: 'List[int]') -> 'bool':
        self.f = [False for i in range(len(nums))]
        self.nums = nums
        return self.do(len(self.nums))
    
    def do(self, lens):
        if lens == 1:
            return True
        if self.f[lens-1]: return False
        i = lens-2
        while i >= 0:
            if lens-1-i<=self.nums[i]:
                flag = self.do(i+1)
                if flag: return True
                self.f[i] = True
            i -= 1
        return False

if __name__ == "__main__":
    print(
        Solution().canJump(
            [2,0,6,9,8,4,5,0,8,9,1,2,9,6,8,8,0,6,3,1,2,2,1,2,6,5,3,1,2,2,6,4,2,4,3,0,0,0,3,8,2,4,0,1,2,0,1,4,6,5,8,0,7,9,3,4,6,6,5,8,9,3,4,3,7,0,4,9,0,9,8,4,3,0,7,7,1,9,1,9,4,9,0,1,9,5,7,7,1,5,8,2,8,2,6,8,2,2,7,5,1,7,9,6]
        )
    )