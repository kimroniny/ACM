class Solution:
    def getPermutation(self, n: 'int', k: 'int') -> 'str':
        fact = [1]
        for i in range(1, n):
            fact.append(fact[-1]*(i+1))
        mylist = list(range(1, n+1))
        yu = -1
        result = []
        while k != 1:
            i = len(mylist)-2
            yu = k % fact[i]
            div = int(k / fact[i])
            if yu == 0:
                div -= 1
                yu = fact[i]
            # print(div)
            result.append(mylist[div])
            k = yu
            mylist.remove(mylist[div])
        result += mylist
        return "".join([str(x) for x in result])

if __name__ == "__main__":
    import math
    n = 1
    for i in range(math.factorial(n)):
        print(Solution().getPermutation(n, i+1))

