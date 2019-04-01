class Solution:
    def setZeroes(self, matrix: 'List[List[int]]') -> 'None':
        """
        Do not return anything, modify matrix in-place instead.
        """
        if len(matrix) == 0: return matrix
        n,m = len(matrix), len(matrix[0])
        flag = matrix[0][0]
        for i in range(n):
            if matrix[i][0] == 0: flag = 0
            for j in range(m):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    if j != 0:  matrix[0][j] = 0
        for i in range(1, n):
            for j in range(1, m):
                if matrix[i][0] == 0 or (matrix[0][j] == 0):
                    matrix[i][j] = 0
        if matrix[0][0] == 0:
            for i in range(m): matrix[0][i] = 0
        if flag: 
            for i in range(n): matrix[i][0] = 0
        return matrix

if __name__ == "__main__":
    print(
        Solution().setZeroes(
            

[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
        )
    )
