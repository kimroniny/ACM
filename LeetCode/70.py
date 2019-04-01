class Solution:
    def climbStairs(self, n: 'int') -> 'int':
        f = [0, 1, 2] + [0 for i in range(n-2)]
        i = 3
        while i <= n:
            f[i] = f[i-1]+f[i-2]
            i += 1
        return f[n]

if __name__ == "__main__":
    print(
        Solution().climbStairs(
            3
        )
    )
