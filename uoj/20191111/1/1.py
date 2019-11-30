m, n, k = (int(x) for x in input().split())

datas = []

for i in range(k):
    datas.append((int(x) for x in input().split()))

res = 0
dp = 

for i in range(k):
    for j in range(k, m):
        