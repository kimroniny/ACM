#include<iostream>
#include<cstdio>
#include<cmath>
using namespace std;


int T, n;
long long dp[100500], ans[100500];

double getLog(int x){
	return log(x)/log(2.0);
} 

long long deal(int x){
	if (dp[x] > 0) return dp[x];
	for (long long i = 2; i <= x/2; i++){
		long long x1, y1, y2, x2, j = x-i;
		x1 = deal(i); x1 = max(x1, i);
		x2 = deal(j); x2 = max(x2, j);
		dp[x] = max(dp[x], x1*x2);
	}
	return dp[x];
}

int main(){
	cin >> T;
	dp[2] = 1;
	dp[3] = 2;
	for (int i = 0; i < T; i++){
		cin >> n;
		cout << "output:" << deal(n) << endl;
	}
	
}

