class Solution:
    def sortColors(self, nums):
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        p = 0
        for i in range(n):
            if nums[i] == 0: nums[p], nums[i] = nums[i], nums[p]; p += 1
        p = n-1
        for i in range(n-1, -1, -1):
            if nums[i] == 2: nums[p], nums[i] = nums[i], nums[p]; p -= 1
        return nums

if __name__ == "__main__":
    print(
        Solution().sortColors(
           [2,2,1,1,0]
        )
    )