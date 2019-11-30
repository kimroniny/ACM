#  
n = int(input())
nums = [int(x) for x in input().split(' ')]

dic = {}

for x in nums:
    dic[x] = (dic[x] + 1) if x in dic else 1
    if dic[x] > n//2: 
        print(x)
        break
