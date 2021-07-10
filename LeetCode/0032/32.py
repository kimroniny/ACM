class Solution:
    def longestValidParentheses(self, s: 'str') -> 'int':
        que = []
        s = list(s)
        for i in range(len(s)):
            if s[i] == ')':
                if len(que) > 0:
                    x = que.pop()
                    s[x] = '.'
                    s[i] = '.'
            elif s[i] == '(':
                que.append(i)
        cnt = 0
        result = 0
        s.append('&')
        for i in range(len(s)):
            if s[i] == '.':
                cnt += 1
            else:
                result = max(result, cnt) if cnt !=0 else result
                cnt = 0
        return result

if __name__ == "__main__":
    print(
        Solution().longestValidParentheses(
            '(()'
        )
    )
        
                    
