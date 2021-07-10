class Solution:
    board = []
    word = ""
    flag = []
    def exist(self, board: 'List[List[str]]', word: 'str') -> 'bool':
        n, m = len(board), len(board[0])
        self.board = board
        self.word = word
        self.flag = [[False for i in range(m)] for i in range(n)]
        loc = []
        for i in range(n):
            for j in range(m):
                c = board[i][j]
                if (len(word)>0 and c == word[0]) or (len(word) == 0 and c == word):
                    loc.append((i, j))
        
        for l, r in loc:
            self.flag = [[False for i in range(m)] for i in range(n)]
            if self.do(l, r, 0, len(word)): 
                print(self.flag)
                return True
        return False
    
    def do(self, l, r, num, length):
        if num >= length: return True
        if l < 0 or l >= len(self.board) or r < 0 or r >= len(self.board[0]) or self.flag[l][r]: return False
        if self.board[l][r] == self.word[num]: 
            self.flag[l][r] = True
            flag = self.do(l+1, r, num+1, length) or self.do(l-1, r, num+1, length) or self.do(l, r+1, num+1, length) or self.do(l, r-1, num+1, length)
            self.flag[l][r] = flag
            return flag
        else: return False

if __name__ == "__main__":
    print(
        Solution().exist(
            [
                ["a","a","a"],
                ["a","b","b"],
                ["a","b","b"],
                ["b","b","b"],
                ["b","b","b"],
                ["a","a","a"],
                ["b","b","b"],
                ["a","b","b"],
                ["a","a","b"],
                ["a","b","a"]
            ],
            "aabaaaabbb"
        )
        

    )