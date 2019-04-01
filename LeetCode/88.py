class Solution:
    def merge(self, nums1: 'List[int]', m: 'int', nums2: 'List[int]', n: 'int') -> 'None':
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i, j = m-1, n-1
        k = len(nums1)-1
        while i >= 0 or j >= 0:
            if (i >= 0 and (j < 0 or (j >=0 and nums1[i] >= nums2[j]))):
                nums1[k] = nums1[i]
                k, i = k-1, i-1
            else:
                nums1[k] = nums2[j]
                k, j = k-1, j-1
        return nums1

if __name__ == "__main__":
    print(
        Solution().merge(
            nums1 = [1,2,3,0,0,0], m = 3, 
            nums2 = [2,5,6],       n = 3
        )
    )
        