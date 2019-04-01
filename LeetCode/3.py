class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """ 
        p = {}  
        f = [0 for c in s] 
        maxlen = -1
        for k, v in enumerate(s):
            if v not in p.keys():
                p[v] = k
                f[k] = f[k-1]+1 if k > 0 else 1
            else:
                f[k] = k-p[v] if k-p[v] <= f[k-1] else f[k-1]+1
                p[v] = k
            maxlen = max(maxlen, f[k])
        # print(f)
        return maxlen

print(Solution().lengthOfLongestSubstring('pwwkew'))


        