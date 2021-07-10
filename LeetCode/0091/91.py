class Solution:
    def numDecodings(self, s: 'str') -> 'int':
        dp = [0] * (len(s))
        for i in range(len(s)):
            last, last1 = 0, 0
            if i > 0:
                x = int(s[i-1:i+1])
                if x >= 10 and x <= 26:  
                    if i >= 2: last = dp[i-2]
                    else: last = 1
                else: last = 0
                last1 = dp[i-1] if s[i] != '0' else 0
            else:
                last1 = 1 if s[i] != '0' else 0

            dp[i] = last1 + last
        return dp[len(s)-1]

if __name__ == "__main__":
    print(
        Solution().numDecodings(
            '01'
        )
    )
         