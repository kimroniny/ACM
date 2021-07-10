class Solution:
    maps = {
            '2':'abc',
            '3':'def',
            '4':'ghi',
            '5':'jkl',
            '6':'mno',
            '7':'pqrs',
            '8':'tuv',
            '9':'wxyz',
        }
    s = ''
    result = []
    def letterCombinations(self, digits: 'str') -> 'List[str]':
        self.result = []
        self.s = ''
        if len(digits) > 0:
            self.do(digits, 0)
        return self.result

    def do(self, string, x: int):
        if x >= len(string):
            self.result.append(self.s)
            return
        if string[x] not in self.maps.keys():
            self.do(string, x+1)
            return
        for c in self.maps[string[x]]:
            self.s += c
            self.do(string, x+1)
            self.s = self.s[0:-1]

if __name__ == "__main__":
    print(
        Solution().letterCombinations(
            '123'
        )
    )
