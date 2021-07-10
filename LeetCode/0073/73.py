class Solution:
    def setZeroes(self, matrix: 'List[List[int]]') -> 'None':
        """
        Do not return anything, modify matrix in-place instead.
        """
        if len(matrix) == 0: return matrix
        n,m = len(matrix), len(matrix[0])
        row = 1
        col = 1
        for i in range(n):
            flag = 0
            for x in matrix[i]:
                if x == 0:
                    flag = 1
                    break
            row = (row << 1) + flag
            # print("row:", row)
        for i in range(m):
            flag = 0
            for j in range(n):
                if matrix[j][i] == 0:
                    flag = 1
                    break
            col = (col << 1) + flag
            # print("col:", col)
        # print(col, row)
        for i in range(n-1, -1, -1):
            if row & 1 == 1:
                for j in range(m):
                    matrix[i][j] = 0
            row = row >> 1
        for j in range(m-1, -1, -1):
            if col & 1 == 1:
                for i in range(n):
                    matrix[i][j] = 0
            col = col >> 1
        # print(col, row)
        return matrix

if __name__ == "__main__":
    print(
        Solution().setZeroes(
            
[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
        )
    )
                
                