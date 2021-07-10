class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        # s[-1:None:-1] reverse
        # 2**31
        intmax = 2147483648
        p = int(x/abs(x))
        s = str(abs(x))
        x = int(''.join(reversed(s)))*p
        if x > intmax-1 or x < -intmax:
            return 0
        return x


if __name__ == "__main__":
    print(Solution().reverse(-120))

        
        