class Solution:
    def generateMatrix(self, n: 'int') -> 'List[List[int]]':
        if n == 0: return []
        matrix = [[1 for i in range(n)] for i in range(n)]
        n, m = len(matrix), len(matrix[0])
        ls, rs, le, re = 0, 0, n-1, m-1
        s = 1
        while ls <= le and rs <= re:
            matrix[ls][rs:re+1] = range(s, s+le-ls+1)
            s += le-ls+1
            i = ls+1
            while i < le: 
                matrix[i][re] = s
                s += 1
                i += 1
            if ls < le: 
                matrix[le][rs:re+1] = list(reversed(range(s, s+re+1-rs)))
                s += re+1-rs
            if rs < re:
                i = le-1
                while i > ls: 
                    matrix[i][rs] = s
                    s += 1
                    i -= 1
            ls, le, rs, re = ls+1, le-1, rs+1, re-1
        return matrix

if __name__ == "__main__":
    print(
        Solution().generateMatrix(1)
    )