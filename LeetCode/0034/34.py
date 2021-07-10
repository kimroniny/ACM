class Solution:
    def searchRange(self, nums: 'List[int]', target: 'int') -> 'List[int]':
        l, r = 0, len(nums)-1
        while l <= r:
            mid = int((l+r)/2)
            if nums[mid] <= target:
                l = mid+1
            else:
                r = mid-1
        maxp = l

        l, r = 0, len(nums)-1
        while l <= r:
            mid = int((l+r)/2)
            if nums[mid] < target:
                l = mid+1
            else:
                r = mid-1
        minp = r
        return [maxp, minp]
    
if __name__ == "__main__":
    print(
        Solution().searchRange(
            [3, 3, 4], 4
        )
    )