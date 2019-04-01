class Solution:
    def mySqrt(self, x: 'int') -> 'int':
        if x in [1, 2, 3]:
            return 1
        if x == 0:
            return 0
        l, r = 2, x-1
        while l <= r:
            mid = int((l+r)/2)
            if mid * mid < x:
                l = mid+1
            elif mid * mid > x:
                r = mid-1
            else:
                return mid
        return r

if __name__ == "__main__":
    import math
    for x in range(100):
        y = Solution().mySqrt(x)
        if int(math.sqrt(x)) != y:
            print(x, y)