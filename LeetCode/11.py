class Solution:
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        ans = 0
        lenh = len(height)
        i, j = 0, lenh-1
        while i < j:
            ans = max(ans, min(height[i],height[j])*(j-i))
            if height[i] < height[j]:
                i = i + 1
            else:
                j = j - 1
        return ans

if __name__ == "__main__":
    print(Solution().maxArea(
        [1,8,6,2,5,4,8,3,7]
    ))
        
            
        