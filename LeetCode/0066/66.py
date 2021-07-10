class Solution:
    def plusOne(self, digits: 'List[int]') -> 'List[int]':
        i = len(digits)-1
        cnt = 1
        while i >= 0:
            if digits[i]+cnt >= 10:
                digits[i], cnt = 0, 1
                i -= 1
            else:
                digits[i] += cnt
                return digits
        if cnt != 0:
            return [1]+digits
        else:
            return digits


if __name__ == "__main__":
    print(
        Solution().plusOne(
            [0]
        )
    )