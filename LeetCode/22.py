class Solution:
    result = ''
    result_list = []
    def generateParenthesis(self, n: 'int') -> 'List[str]':
        add = n
        dec = n
        self.result = ''
        self.result_list = []
        self.do(add, dec)
        return self.result_list

    def do(self, add, dec):
        if add > 0:
            self.result = self.result + '('
            self.do(add-1, dec)
            self.result = self.result[0:-1]
        if dec > 0 and dec > add:
            self.result = self.result + ')'
            self.do(add, dec-1)
            self.result = self.result[0:-1]
        if add == 0 and dec == 0:
            self.result_list.append(self.result)


if __name__ == "__main__":
    print(
        Solution().generateParenthesis(3)
    )

        