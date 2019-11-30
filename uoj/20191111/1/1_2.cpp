#include<iostream>
#include<cstdio>

using namespace std;

typedef struct{
	int a, b, c;
}Info;

Info e[1005];
int dp[1005][1005]; 

int main(){
	int m, n, k;
	scanf("%d%d%d", &m, &n, &k);
	for (int i = 0; i < k; i++){
		scanf("%d%d%d", &e[i].a, &e[i].b, &e[i].c);	
	}
	for (int i = 0; i < k; i++){
		for (int j = m; j >= 0; j--)
		if (j >= e[i].a){
			for (int k = n; k >= 0; k--)
			if (k >= e[i].b){
				dp[j][k] = max(dp[j][k], dp[j-e[i].a][k-e[i].b]+e[i].c);
			}
		}
	}
	int ans = 0;
	for (int i = 0; i <= m; i++)
		for (int j = 0; j <= n; j++)
			ans = max(ans, dp[i][j]);
	cout << ans;
} 
