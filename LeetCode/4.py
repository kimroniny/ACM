class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        return self.deal(nums1, nums2)*1.0
        
    def deal(self, nums1, nums2):
        while True:
            len1, len2 = len(nums1), len(nums2)
            if len1 > len2:
                len2, len1 = len1, len2
                nums2, nums1 = nums1, nums2
            if len1 == 0:
                return nums2[int(len2/2)] if len2%2!=0 else (nums2[int(len2/2)-1]+nums2[int(len2/2)])/2
            if len2 == 0:
                 return nums1[int(len1/2)] if len1%2!=0 else (nums1[int(len1/2)-1]+nums1[int(len1/2)])/2
            mid1 = int((len1-1)/2)
            mid2 = int((len1+len2-1)/2)-(mid1+1) # int((len2)/2) number of A and B is not equal.
            if mid2+1>=len2:
                nums2, nums1 = nums1, nums2
                len2, len1 = len1, len2
                mid2,mid1 = mid1, mid2
            if mid1+1>=len1:
                if mid2+1<len2:
                    if nums1[mid1] < nums2[mid2+1]:
                        if mid2 >= 0:
                            return max(nums1[mid1], nums2[mid2]) if (len1+len2)%2!=0 else (max(nums1[mid1], nums2[mid2])+nums2[mid2+1])/2
                        else:
                            return nums1[mid1] if (len1+len2)%2!=0 else (nums1[mid1]+nums2[mid2+1])/2
                    else:
                        return nums2[mid2+1] if (len1+len2)%2!=0 else (nums2[mid2+1]+nums1[mid1])/2
                else:
                    return (nums1[mid1]+nums2[mid2])/2

            if nums1[mid1] <= nums2[mid2+1] and nums1[mid1+1] >= nums2[mid2]:
                return min(nums1[mid1+1], nums2[mid2+1]) if (len1+len2)%2!=0 else (max(nums1[mid1], nums2[mid2]) + min(nums1[mid1+1], nums2[mid2+1]))/2
            elif nums1[mid1] > nums2[mid2+1]:
                nums1, nums2 = nums1[:mid1+1], nums2[mid2+1:]
            elif nums1[mid1+1] < nums2[mid2]:
                nums1, nums2 = nums1[mid1+1:], nums2[:mid2+1]

if __name__ == "__main__":
    solution = Solution()
    # nums1 = [1, 3, 4]
    # nums2 = [2]
    # nums1 = [1]
    # nums2 = [2, 3, 4]
    # nums1 = [1, 2]
    # nums2 = [3, 4]
    nums1 = [4]
    nums2 = [1, 2, 3]
    result = solution.findMedianSortedArrays(nums1, nums2)
    print(result)


'''
1 2 3 4 5
6 7 8 9
5

3 4 5 
6 7
5


'''