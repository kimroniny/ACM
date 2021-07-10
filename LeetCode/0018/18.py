class Solution:
    def fourSum(self, nums: 'List[int]', target: 'int') -> 'List[List[int]]':
        nums = sorted(nums)
        lens = len(nums)
        result = []
        for i in range(lens):
            j = i + 1
            while j < lens:
                p, q = j+1, lens-1
                x = target-nums[i]-nums[j]
                while p < q:
                    sums = nums[p]+nums[q]
                    if sums > x:
                        q-=1
                    elif sums < x:
                        p += 1
                    else:
                        k = [nums[i], nums[j], nums[p], nums[q]]
                        if k not in result:
                            result.append(k)
                        p += 1
                        q -= 1
                j += 1
        return result

if __name__ == "__main__":
    print(
        Solution().fourSum(
            [1, 0, -1, 0, -2, 2],
            0
        )
    )