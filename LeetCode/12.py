class Solution:
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        ans = ''
        roman = {1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}
        keys = list(roman.keys())
    
        i = len(keys)-1
        while i > -1:
            x = int(num / keys[i])
            y = num % keys[i]
            if (i-1>=0 and keys[i] in [5, 50, 500] and keys[i]-keys[i-1]<=num and num < keys[i]):
                ans = ans + roman[keys[i-1]] + roman[keys[i]]
                num = num-keys[i]+keys[i-1]
            elif (i-2>=0 and keys[i] in [10, 100, 1000] and keys[i]-keys[i-2]<=num and num < keys[i]):
                ans = ans + roman[keys[i-2]] + roman[keys[i]]
                num = num-keys[i]+keys[i-2]
            else:
                num = y
                ans = ans + roman[keys[i]]*x
            if x == 0:
                i = i - 1
            if num == 0:
                break
        return ans

if __name__ == "__main__":
    print(Solution().intToRoman(1994))
                        
