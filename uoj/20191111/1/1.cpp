#include<iostream>
#include<cstdio>
#include<cstring>

using namespace std;

typedef struct{
	int a, b, c;
}Info;
typedef struct{
	int a, b;
}Info2;

Info e[1005];
Info2 dp1[1005],dp2[1005];

int main(){
	int m, n, k;
	scanf("%d%d%d", &m, &n, &k);
	for (int i = 0; i < k; i++){
		scanf("%d%d%d", &e[i].a, &e[i].b, &e[i].c);	
	}
	for (int i = 0; i < k; i++){
		for (int j = m; j >= e[i].a; j--){
			if (dp1[j-e[i].a].b + e[i].c > dp1[j].b){
				if (dp1[j-e[i].a].a + e[i].b <= n){
					dp1[j].b = dp1[j-e[i].a].b + e[i].c;
					dp1[j].a = dp1[j-e[i].a].a + e[i].b;
				} 
			}else if (dp1[j-e[i].a].b + e[i].c == dp1[j].b){
				if (dp1[j-e[i].a].a + e[i].b < dp1[j].a){
					dp1[j].a = dp1[j-e[i].a].a + e[i].b;
				}
			}
		}
	}
	for (int i = 0; i < k; i++){
		for (int j = n; j >= e[i].b; j--){
			if (dp2[j-e[i].b].b + e[i].c > dp2[j].b){
				if (dp2[j-e[i].b].a + e[i].a <= m){
					dp2[j].b = dp2[j-e[i].b].b + e[i].c;
					dp2[j].a = dp2[j-e[i].b].a + e[i].a;
				} 
			}else if (dp2[j-e[i].b].b + e[i].c == dp2[j].b){
				if (dp2[j-e[i].b].a + e[i].a < dp2[j].a){
					dp2[j].a = dp2[j-e[i].b].a + e[i].a;
				}
			}
		}
	}
	printf("%d\n", max(dp1[m].b, dp2[n].b));
	
	return 0;
}
