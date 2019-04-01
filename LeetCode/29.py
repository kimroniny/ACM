class Solution:
    def divide(self, dividend: 'int', divisor: 'int') -> 'int':
        sign = 1 if (dividend >= 0 and divisor >= 0) or (dividend < 0 and divisor < 0) else -1
        dividend, divisor = abs(dividend), abs(divisor)
        maxint = 2**31-1
        result = 0
        divisor2 = divisor
        cnt = 1
        while dividend >= divisor:
            while dividend < divisor2:
                divisor2 -= divisor
                cnt -= 1
            dividend -= divisor2
            result += cnt
            cnt += 1
            divisor2 += divisor
        if sign == -1 and result >= maxint+1:
            return -(maxint+1)
        elif sign == 1 and result >= maxint:
            return maxint
        else:
            return sign*result

if __name__ == "__main__":
    print(
        Solution().divide(
            10, 2
        )
    )
