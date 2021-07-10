class Solution:
    def removeDuplicates(self, nums) -> int:
        p, cnt = 0, 1
        l = len(nums)
        if l == 0: return 0
        for i in range(1, l):
            if nums[i] != nums[p]:
                p += 1
                nums[i], nums[p] = nums[p], nums[i]
                cnt = 1
            elif cnt < 2:
                p += 1
                nums[i], nums[p] = nums[p], nums[i]
                cnt += 1
        return p + 1

if __name__ == "__main__":
    print(
        Solution().removeDuplicates(
            []
        )
    )