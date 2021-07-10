class Solution:
    def search(self, nums: 'List[int]', target: 'int') -> 'int':
        l, r = 0, len(nums)-1
        while l <= r:
            mid = int((l+r)/2)
            x = nums[mid]
            if x == target: return True
            if nums[l] <= x:
                if target >= nums[l] and target < x:
                    r = mid-1
                else:
                    l = mid+1
            else:
                if target > x and target <= nums[r]:
                    l = mid+1
                else:
                    r = mid-1
        return False

if __name__ == "__main__":
    print(
        Solution().search(
            [4,5,6,7,0,1,2], 8
        )
    )
            
