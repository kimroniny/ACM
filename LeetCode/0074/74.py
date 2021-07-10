class Solution:
    def searchMatrix(self, matrix, target: int) -> bool:
        for mylist in matrix:
            if target in mylist:
                return True
        return False

if __name__ == "__main__":
    print(
        Solution().searchMatrix(
             [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
], 3
        )
    )