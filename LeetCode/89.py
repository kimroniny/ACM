# 4
# 0000  0   1  0
# 0001  1   2  1
# 0011  3   3  3
# 0010  2   4  2
# 0110  6   5
# 0111  7   6
# 0101  5   7
# 0100  4   8  
# 1100

class Solution:
    result = None
    def grayCode(self, n: int):
        self.result = [0]*(1<<n)
        self.deal(0, 0, n)
        return self.result
    
    def deal(self, loc, lev, n):
        p = lev+1
        while p <= n:
            nloc = (1<<p)-loc-1
            self.result[nloc] = self.result[loc] + (1 << (p-1))
            self.deal(nloc, p, n)
            p += 1

if __name__ == "__main__":
    print(
        Solution().grayCode(4)
    )
        