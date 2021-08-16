#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>

typedef long long LL;
int p,q;
int MAX = 1<<31-1, MOD = 1000000007;
LL d[51][51][55];

LL dfs(int x, int y, int rest) {
    if (x < 0 || x >= p || y < 0 || y >= q) return 1;
    if (rest == 0) return 0;
    if (d[x][y][rest] != MAX) return d[x][y][rest];
    d[x][y][rest] = 0;
    d[x][y][rest] += 
        dfs(x+1,y,rest-1) + 
        dfs(x, y+1, rest-1) + 
        dfs(x-1,y,rest-1) + 
        dfs(x, y-1, rest-1);
    d[x][y][rest] %= MOD;
    return d[x][y][rest];
}

int findPaths(int m, int n, int maxMove, int startRow, int startColumn){
    p = m; q = n;
    int i, j, k;
    for (i = 0; i < m; i++)
        for (j = 0; j < n; j++) 
            for (k = 0; k < maxMove+1; k++)
                d[i][j][k] = MAX;
    return dfs(startRow, startColumn, maxMove);
}

int main(void) {
    printf("%d", findPaths(
        1,3,3,0,1
    ));
    return 0;
}