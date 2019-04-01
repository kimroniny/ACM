class Solution:
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        # regex: x=re.search('^\s*([-\+]?\d+)', str)
        maxint = 2147483648
        try:
            a = str.strip(' ')
            p = 1
            if a[0] in '-+':
                if a[0] == '-':
                    p = -1
                a = a[1:]
            # print(a)
            i = 0
            while i < len(a) and a[i] in '1234567890':
                i = i + 1
            ints = int(a[0:i])*p
            
            if ints >= maxint:
                return maxint-1
            if ints < -maxint:
                return  -maxint
            return ints
        except Exception as e:
            # print(e)
            return 0
        
        

if __name__ == "__main__":
    print(Solution().myAtoi('-45'))