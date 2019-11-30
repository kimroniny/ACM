n = int(input())

for i in range(n):
    input()
    s, ss = input().split()
    dp = [[0]*len(ss) for p in range(len(s))]
    dp[0][0] = 1 if s[0] == ss[0] else 0
    for j in range(1, len(s)): dp[j][0] = 1 if s[j] == ss[0] else dp[j-1][0]
    for j in range(1, len(ss)): dp[0][j] = 1 if s[0] == ss[j] else dp[0][j-1]
    for p in range(1, len(s)):
        for q in range(1, len(ss)):
            dp[p][q] = max(dp[p-1][q], dp[p][q-1]) if s[p] != ss[q] else dp[p-1][q-1]+1
    print(dp[-1][-1])
            
        
