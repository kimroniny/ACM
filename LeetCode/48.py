class Solution:
    def rotate(self, matrix: 'List[List[int]]') -> 'None':
        """
        Do not return anything, modify matrix in-place instead.
        """
        lens = len(matrix)
        half = (lens-1)/2.0
        for i in range(int(lens/2)):
            for j in range(i, lens-i-1):
                p = 4
                ti, tj, tval = i, j, matrix[i][j]
                while p > 0:
                    i1, j1 = tj, half*2-ti
                    ti, tj, tval, matrix[i1][j1] = i1, j1, matrix[i1][j1], tval
                    p -= 1
        return matrix

if __name__ == "__main__":
    print(
        Solution().rotate(
            [
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
]
        )
    )
