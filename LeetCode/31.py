class Solution:
    def nextPermutation(self, nums: 'List[int]') -> 'None':
        """
        Do not return anything, modify nums in-place instead.
        """
        zengi, zengj, jiani, jianj = -1, -1, -1, -1
        flag = False
        for i in range(1, len(nums)):
            if nums[i-1] < nums[i]:
                zengi, zengj = i-1, i
                flag = False
            elif nums[i-1] > nums[i]:
                if not flag:
                    jiani = i-1
                flag = True
                jianj = i
        if jiani < zengi:
            nums[jiani], nums[jianj] = nums[jianj], nums[jiani]
        else:
            i, j = jiani, jianj
            while i < j:
                nums[i], nums[j] = nums[j], nums[i]
                i, j = i+1, j-1
            if jiani > 0:
                nums[jiani], nums[jiani-1] = nums[jiani-1], nums[jiani]
        
        return nums
    
if __name__ == "__main__":
    print(
        Solution().nextPermutation(
            [1, 2, 3]
        )
    )
                