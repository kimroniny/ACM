class Solution:
    def isValid(self, s: 'str') -> 'bool':
        maps = {
            '(': 1,
            ')': 10,
            '{': 2,
            '}': 20,
            '[': 3,
            ']': 30,
        }
        que = [0 for i in range(len(s))]
        p = -1
        for c in s:
            if maps[c] < 10:
                p += 1
                que[p] = maps[c]
            else:
                if que[p] == int(maps[c]/10):
                    p -= 1
                else:
                    return False
        
        return True if p == -1 else False

if __name__ == "__main__":
    print(
        Solution().isValid(
        '{[]}'
        )
    )