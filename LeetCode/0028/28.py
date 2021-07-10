class Solution:
    def strStr(self, haystack: 'str', needle: 'str') -> 'int':
        lens2 = len(needle)
        for i in range(len(haystack)+1):
            if haystack[i:i+lens2] == needle:
                return i
        return -1

if __name__ == "__main__":
    print(
        Solution().strStr(
            'hello',
            'll'
        )
    )
        