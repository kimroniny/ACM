class Solution:
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        l1, l2 = len(num1), len(num2)
        i, j = l1-1, l2-1
        result = [0 for i in range(l1+l2)]
        while i >= 0:
            jin = 0
            j = l2-1
            while j >= 0:
                multi = int(num1[i])*int(num2[j])+result[l1-1-i+l2-1-j]+jin
                jin = multi // 10
                result[l1-1-i+l2-1-j] = multi % 10
                j -= 1
            result[l1-1-i+l2-1-j] = jin
            i -= 1
        i = len(result)-1
        while i >= 0 and result[i] == 0:
            i -= 1
        return (''.join(reversed([str(x) for x in result[0:i+1]])) if i > -1 else '0')

if __name__ == "__main__":
    for i in range(10000):
        for j in range(10000):
            result = Solution().multiply(str(i), str(j))
            if str(i*j) != result:
                print([i, j, result])



                