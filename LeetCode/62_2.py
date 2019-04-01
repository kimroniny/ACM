class Solution:
    def uniquePaths(self, m: 'int', n: 'int') -> 'int':
        a = [1]*m
        for i in range(1, n):
            for j in range(1, m):
                a[j] = a[j-1]+a[j]
        return a[m-1]

if __name__ == "__main__":
    print(
        Solution().uniquePaths(
            7, 3
        )
    )