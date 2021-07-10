class Solution:
    def minPathSum(self, grid: 'List[List[int]]'):
        dp = [[0]*len(grid) for i in range(len(grid[0]))]
        for p, x in enumerate(grid[0]):
            if p == 0: dp[0][p] = x
            else: dp[0][p] = dp[0][p-1]+x
        for i in range(len(grid)):
            if i == 0: dp[i][0] = grid[i][0]
            else: dp[i][0] = dp[i-1][0]+grid[i][0]
        for i in range(1, len(grid)):
            for j in range(1, len(grid[0])):
                dp[i][j] = min(dp[i-1][j], dp[i][j-1])+grid[i][j]
        return dp[-1][-1]

if __name__ == "__main__":
    print(
        Solution().minPathSum(
            [
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
        )
    )