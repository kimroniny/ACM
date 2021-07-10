class Solution(object):
    def divide(self, dividend, divisor):
        """
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        
        if dividend == 0:
            return 0
        if dividend < 0 and divisor < 0:
            return min(2147483647, self.helper(-dividend, -divisor))
        if dividend < 0 and divisor > 0:
            return max(-2147483648, -self.helper(-dividend, divisor))
        if dividend > 0 and divisor < 0:
            return max(-2147483648, -self.helper(dividend, -divisor))
        return min(2147483647, self.helper(dividend, divisor))
        
    def helper(self, dividend, divisor):
        if divisor > dividend:
            return 0
        
        div = divisor
        res = 1
        while div + div <= dividend:
            div += div
            res += res

        return res + self.helper(dividend - div, divisor)