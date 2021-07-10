class Solution:
    def searchInsert(self, nums: 'List[int]', target: 'int') -> 'int':
        lens = len(nums)
        l, r = 0, lens-1
        while l <= r:
            mid = int((l+r)/2)
            if nums[mid] < target:
                l = mid+1
            else:
                r = mid-1
        # if l-1>=0 and nums[l] 
        return l

if __name__ == "__main__":
    print(
        Solution().searchInsert(
            [1, 2, 3, 4, 5], 2.5
        )
    )