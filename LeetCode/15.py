class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ans = []
        nums = sorted(nums)
        lens, i = len(nums), 0
        while i < lens-2:
            cur = nums[i]
            if cur > 0:
                break
            if i > 0 and cur == nums[i-1]:
                i = i + 1
                continue
            j, k = i+1, lens-1
            while j < k:
                if nums[j] + nums[k] > -cur:
                    k -= 1
                elif nums[j] + nums[k] < -cur:
                    j += 1
                else:
                    ans.append([cur, nums[j], nums[k]])
                    j += 1
                    while (j < k and nums[j] == nums[j-1]):
                        j+=1
                    k -= 1
                    while (j < k and nums[k] == nums[k+1]):
                        k -= 1
            i = i+1
        return ans
        

if __name__ == "__main__":
    print(Solution().threeSum(
        [0, 0,0]
    ))
        