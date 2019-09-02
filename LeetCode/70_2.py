class Solution:
    def climbStairs(self, n: 'int') -> 'int':
        a, b = 1, 2
        if n == 1: return a
        if n == 2: return b
        if n == 0: return 0
        i = 3
        while i <= n:
            c = a+b
            b, a = c, b
            i += 1
        return b

if __name__ == "__main__":
    print(
        Solution().climbStairs(
            4
        )
    )
