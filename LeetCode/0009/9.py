class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        # s[::-1]
        # if str(x) == ''.join(reversed(str(x))):
        if str(x) == (str(x))[-1:None:-1]:
            return True
        return False

if __name__ == "__main__":
    print(Solution().isPalindrome(101))