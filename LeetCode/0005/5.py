class Solution:
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        lens = len(s)
        i = 0
        maxlen = -1
        maxstr = ""
        while i < lens:
            c = s[i]
            j, k = i+1, i - 1
            while j < lens and s[j] == c:
                j = j+1
            while k > -1 and s[k] == c:
                k = k - 1
            i = j
            p = j-k-1
            while j < lens and k > -1 and s[j] == s[k]:
                j, k, p = j + 1, k - 1, p + 2
            if maxlen < p:
                maxlen, maxstr = p, s[k+1:j]

        return maxstr

if __name__ == "__main__":
    
    string = "abbc"
    s = Solution()
    print(s.longestPalindrome(string))

            