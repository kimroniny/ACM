class Solution:
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        maps = {
            'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000
        }
        alts = {
            'I': ['V', 'X'], 'X': ['L', 'C'], 'C': ['D', 'M']
        }
        lens, i, ans = len(s), 0, 0
        while i < lens:
            if s[i] in alts.keys():
                if (i + 1 < lens and s[i+1] in alts[s[i]]):
                    ans += maps[s[i+1]]-maps[s[i]]
                    i += 2
                    continue
                else:
                    ans += maps[s[i]]
            else:
                ans += maps[s[i]]
            i += 1
        return ans

if __name__ == "__main__":
    print(Solution().romanToInt('MCMXCIV'))