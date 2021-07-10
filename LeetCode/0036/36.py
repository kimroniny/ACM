class Solution:
    def isValidSudoku(self, board: 'List[List[str]]') -> 'bool':
        boardr = [[] for i in range(9)]
        boarddiv = [[] for i in range(9)]
        for i in range(9):
            for j, col in enumerate(board):
                boardr[i].append(col[i])
                boarddiv[int(i/3)+int(j/3)*3].append(col[i])
        for i in range(9):
            for j in range(9):
                x = board[i][j]
                if x != '.':
                    if board[i].count(x)>1 or boardr[j].count(x) > 1 or boarddiv[int(j/3)+int(i/3)*3].count(x) > 1:
                        # print(i, j, x)
                        return False
        return True

if __name__ == "__main__":
    board = [[] for i in range(9)]
    for i in range(9):
        for j in range(9):
            board[i].append(i*9+j)
    

    print(
        Solution().isValidSudoku(
           [
               ["7",".",".",".","4",".",".",".","."],
               [".",".",".","8","6","5",".",".","."],
               [".","1",".","2",".",".",".",".","."],
               [".",".",".",".",".","9",".",".","."],
               [".",".",".",".","5",".","5",".","."],
               [".",".",".",".",".",".",".",".","."],
               [".",".",".",".",".",".","2",".","."],
               [".",".",".",".",".",".",".",".","."],
               [".",".",".",".",".",".",".",".","."]
            ]
        )
    )

        
        
            