class Solution:
    n = 0
    k = 0
    result = []
    result_all = []
    def combine(self, n: int, k: int):
        if n == 0 or k == 0 or k > n: return []
        self.n, self.k = n, k
        self.result, self.result_all = [], []
        self.do(1, 1)
        return self.result_all

    def do(self, x, k):
        if k > self.k: 
            self.result_all.append([x for x in self.result])
            return
        while x <= self.n-self.k+k:
            self.result.append(x)
            self.do(x+1, k+1)
            x += 1 
            self.result.pop()

if __name__ == "__main__":
    print(
        Solution().combine(
            5, 0
        )
    )

