import queue
import heapq

class  P():
    def __init__(self,a,b):
        self.a = a
        self.b = b
    def __lt__(self, other):
        if self.a<other.a:
            return True
        else:
            return False
    def p(self):
        print(self.a, self.b)

class Solution:
    def eliminateMaximum(self, dist, speed) -> int:
        h = []
        ans = 0
        for v, k in enumerate(dist):
            heapq.heappush(h, P(k, speed[v]))
        while len(h) > 0:
            x = heapq.heappop(h)
            if x.a == 0:
                return ans
            

if __name__ == "__main__":
    # print()
    Solution().eliminateMaximum(
        [3,2,4],
        [5,3,2]
    )