# 计算C，数据太大
class Solution:
    def uniquePaths(self, m: 'int', n: 'int') -> 'int':
        p = m+n-2
        q = m-1 if m-1 < n-1 else n-1
        return self.C(p, q)
    
    def C(self, p, q):
        if q == 0:
            return 1
        zi = p
        mu = q
        for i in range(0, q-1): zi, mu = zi*(zi-1), mu * (mu-1)
        return int(zi/mu)

if __name__ == "__main__":
    print(
        Solution().uniquePaths(
            7, 3
        )
    )
