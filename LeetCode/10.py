class Solution:
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        i, j = 0, 0
        c = ''
        while i < len(s) and j < len(p):
            if p[j] == '.' or (s[i]==p[j] and p[j] != '*'):
                c = p[j]
                j = j + 1
                i = i + 1
                # print(c)
            elif p[j] == '*':
                while i < len(s) and (c == '.' or s[i] == c):
                    i = i + 1
                j = j + 1
            else:
                j = j + 1
                while (j<len(p) and p[j] == '*'):
                    j = j + 1
        
        while j < len(p):
            if p[j] != '*':
                return False
            j = j+1
        return i == len(s) and j == len(p)

if __name__ == "__main__":
    print(Solution().isMatch(
        'mississippi',
        'mis*is*p*.'
        )
    )
                

        