class Solution:
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        p = strs[0]
        for k in strs[1:]:
            # print(k, p)
            p = self.find(k, p)
            if p == '':
                return ''
        return p
    def find(self, s1:str, s2:str):
        lens1, lens2 = len(s1), len(s2)
        i, j = 0, 0
        while i < lens1 and j < lens2:
            if s1[i] == s2[j]:
                i, j = i+1, j+1
            else:
                break
        return s1[0:i]

if __name__ == "__main__":
    print(Solution().longestCommonPrefix(
        ["dog","racecar","car"]
    ))
        