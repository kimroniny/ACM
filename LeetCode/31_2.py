class Solution:
    def nextPermutation(self, nums: 'List[int]') -> 'None':
        """
        Do not return anything, modify nums in-place instead.
        """
        lens = len(nums)
        if lens < 2:
            return nums
        i = lens-1
        if nums[i-1] < nums[i]:
            nums[i-1], nums[i] = nums[i], nums[i-1]
            return nums
        while i > 0 and nums[i-1] >= nums[i]:
            i -= 1
        j = i
        p, q = i, lens-1
        while p < q:
            nums[p], nums[q] = nums[q], nums[p]
            p, q = p+1, q-1
        if j > 0:
            p = j
            while p < lens and nums[j-1] >= nums[p]:
                p += 1
            nums[j-1], nums[p] = nums[p], nums[j-1]
            while p < lens-1 and nums[p] > nums[p+1]:
                nums[p], nums[p+1] = nums[p+1], nums[p]
                p += 1
        
        
        return nums

if __name__ == "__main__":
    print(
        Solution().nextPermutation(
            [1, 2, 1]
        )
    )