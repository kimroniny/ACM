class Solution:
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        p = 1
        s = '1'
        # results = ''
        while p < n:
            i = 1
            cnt = 1
            news = ''
            s += '%'
            while i < len(s):
                if s[i] != s[i-1]:
                    news += str(cnt)+s[i-1]
                    cnt = 1
                else:
                    cnt += 1
                i+=1
            s = news
            p += 1
        return s

if __name__ == "__main__":
    print(
        Solution().countAndSay(6)
    )
            