重点在于转化。

回文串的特点在于对称相等，即反转过来和原串是一样的，所以可以转化为求：两个字符串的最长公共子序列。

```c
int longestPalindromeSubseq(char * s){
    int i, j, n = strlen(s)+1;
    int d[n][n];
    memset(d,0,sizeof(d)); 
    for (i = 1; i < n; i++) {
        for (j = n-2; j>=0; j--) {
            if (s[i-1] == s[j]) {
                d[i][j] = d[i-1][j+1] + 1;
            }else{
                d[i][j] = d[i-1][j] > d[i][j+1]?d[i-1][j]:d[i][j+1];
            }
        }
    }
    return d[n-1][0];
}
```