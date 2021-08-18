#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>

const int MOD = 1000000007;

long long f[100050][4][2];
int checkRecord(int n){
    memset(f, 0, sizeof(f));
    f[0][0][0] = 1;
	f[0][1][0] = 1;
	f[0][2][0] = 1;
    for (int i = 1; i < n; i++) {
		// 'P'
		f[i][0][0] += (f[i-1][0][0] + f[i-1][2][0] + f[i-1][3][0]) % MOD;
		f[i][0][1] += (f[i-1][1][0] + f[i-1][0][1] + f[i-1][2][1] + f[i-1][3][1]) % MOD;
		// 'A'
		f[i][1][0] += (f[i-1][0][0] + f[i-1][2][0] + f[i-1][3][0]) % MOD;
		f[i][1][1] = 0;
		// 'XL'
		f[i][2][0] += f[i-1][0][0] % MOD;
		f[i][2][1] += (f[i-1][0][1] + f[i-1][1][0]) % MOD;
		// 'XLL'
		if (i > 1) {
			f[i][3][0] += f[i-2][0][0] % MOD;
			f[i][3][1] += (f[i-2][0][1] + f[i-2][1][0]) % MOD;
		} else {
			f[i][3][0] = 1;
            f[i][3][1] = 0;
		}
	}
	long long ans = 0;
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < 2; j++) {
			ans += f[n-1][i][j];
			ans %= MOD;
		}
	}
	return ans;
}

void main(void) {
    printf("%d\n", checkRecord(10101));
}