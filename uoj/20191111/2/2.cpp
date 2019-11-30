#include<iostream>
#include<cstdio>
#include<vector>

using namespace std;
int dp[1005][1005];

int main(){
	int n;
	cin >> n;
	for (int i = 0; i < n; i++){
		int l1, l2;
		scanf("%d%d", &l1, &l2);
		char s1[l1], s2[l2];
		scanf("%s%s", s1, s2);
		dp[0][0] = s1[0]==s2[0]?1:0;
		for (int j = 1; j < l1; j++) dp[j][0] = s1[j]==s2[0]?1:dp[j-1][0];
		for (int j = 1; j < l2; j++) dp[0][j] = s1[0]==s2[j]?1:dp[0][j-1];
		for (int p = 1; p < l1; p++){
			for (int q = 1; q < l2; q++){
				dp[p][q] = s1[p]==s2[q]?dp[p-1][q-1]+1:max(dp[p-1][q], dp[p][q-1]);
			}
		}
		printf("%d\n", dp[l1-1][l2-1]);
	}
} 
