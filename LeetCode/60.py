class Solution:
    result = []
    k = 0
    target = 0
    def getPermutation(self, n: 'int', k: 'int') -> 'str':
        self.result = []
        self.target = k
        self.k = 0
        mylist = list(range(1, n+1))
        return self.do(mylist)
        

    
    def do(self, mylist):
        # print("mylist ", mylist)
        if len(mylist) == 0:
            self.k += 1
            if self.k == self.target: return "".join([str(x) for x in self.result])
        for x in mylist:
            self.result.append(x)
            copy = mylist.copy()
            copy.remove(x)
            s = self.do(copy)
            if s is not None: return s
            self.result.pop()


if __name__ == "__main__":
    import math
    n = 4
    for i in range(math.factorial(n)):
        print(Solution().getPermutation(n, i+1))