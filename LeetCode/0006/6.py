class Solution:
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        row = numRows
        if row == 1:
            return s
        a = int(len(s)/(row*2-2))
        b = len(s)%(row*2-2)
        group_num = row*2-2
        group_col = row-1
        col = a*group_col + ((1 + b-row) if b > row else 1)
        result = [[' ' for i in range(col)] for j in range(row)]
        # print(col, row, group_num, group_col)
        for p, c in enumerate(s):
            p = p + 1
            group = int(p/group_num)
            pcol = group * group_col
            rank = p%group_num
            if rank >= row:
                pcol = pcol + 1 + rank-row
                prow = row-(rank-row)
            else:
                pcol = pcol + (1 if rank != 0 else 0)
                prow = rank if rank != 0 else 2 
            # print(prow, pcol, c)
            result[prow-1][pcol-1] = c
        s = ''
        for i in result:
            s = s + ''.join(i)
        return s.replace(' ', '')
            

if __name__ == "__main__":
    s = "PAYPALISHIRING"
    num = 3
    print(Solution().convert(s, num))
    