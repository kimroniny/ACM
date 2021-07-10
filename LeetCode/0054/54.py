# 分治，每个圈看成一个子问题
class Solution:
    def spiralOrder(self, matrix: 'List[List[int]]') -> 'List[int]':
        if len(matrix) == 0: return []
        n, m = len(matrix), len(matrix[0])
        result = []
        ls, rs, le, re = 0, 0, n-1, m-1
        while ls <= le and rs <= re:
            result += matrix[ls][rs:re+1]
            i = ls+1
            while i < le: 
                result.append(matrix[i][re])
                i += 1
            if ls < le: result += list(reversed(matrix[le][rs:re+1]))
            if rs < re:
                i = le-1
                while i > ls: 
                    result.append(matrix[i][rs])
                    i -= 1
            ls, le, rs, re = ls+1, le-1, rs+1, re-1
        return result


if __name__ == "__main__":
    print(
        Solution().spiralOrder(
            [
  
  
]
        )
    )