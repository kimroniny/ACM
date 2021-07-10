class Solution:
    def groupAnagrams(self, strs: 'List[str]') -> 'List[List[str]]':
        maps = self.getmaps()
        result = {}
        for s in strs:
            maps = self.getmaps()
            for c in s:
                maps[ord(c)-ord('a')] += 1
            key = ''
            for i in range(len(maps)):
                key += maps[i]*str(chr(i+ord('a')))
            if key not in result:
                result[key] = [s]
            else:
                result[key].append(s)
        # print(result)
        return list(result.values())

    def getmaps(self):
        return [0 for i in range(26)]

if __name__ == "__main__":
    print(
        Solution().groupAnagrams(
            ["eat", "tea", "tan", "ate", "nat", "bat"]
        )
    )
        