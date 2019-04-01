class Solution:
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums = sorted(nums)
        l = len(nums)
        closest = 2**31-1
        result = 0
        for i in range(l):
            # print(i)
            j, k = i+1, l-1
            x = target-nums[i]
            while j < k:
                sums = nums[j]+nums[k]
                if closest > abs(sums-x):
                    closest = abs(sums-x)
                    result = sums+nums[i]
                if sums < x:
                    j += 1
                elif sums > x:
                    k -= 1
                else:
                    return target
        return result

if __name__ == "__main__":
    result = Solution().threeSumClosest(
        [1,2,4,8,16,32,64,128],
        82
    )
    print(result)