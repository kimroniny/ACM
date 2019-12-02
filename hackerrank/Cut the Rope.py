T = int(input())
dp = [0] * 100500
dp[2] = 1
dp[3] = 2

def deal(x):
    if dp[x] > 0: return dp[x]
    for i in range(2, x//2+1):
        dp[x] = max(dp[x], max(deal(i), i) * max(deal(x-i), x-i))
    return dp[x]

while T > 0:
    T -= 1
    n = int(input())
    # print("output: "+str(deal(n) % 1000))
    print(deal(n)%1000)
    
