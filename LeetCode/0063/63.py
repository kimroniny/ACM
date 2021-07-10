class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: 'List[List[int]]') -> 'int':
        # if obstacleGrid[-1][-1] == 1:
        #     return False
        a = [1-x for x in obstacleGrid[0]]
        for i in range(1, len(a)): a[i] = a[i-1]&a[i]
        # print(a)
        n = len(obstacleGrid)
        m = len(obstacleGrid[0])
        for i in range(1,n):
            if obstacleGrid[i][0] == 1: a[0] = 0
            for j in range(1,m):
                if obstacleGrid[i][j] != 1:
                    a[j] = a[j-1]+a[j]
                else:
                    a[j] = 0
        return a[m-1]

if __name__ == "__main__":
    print(
        Solution().uniquePathsWithObstacles(
            
[[0,1,0,0,0],[1,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0]]
        )
    )
