#include<iostream>
#include<cstdio>
#include<vector>

using namespace std;

vector<int> a; 

int main(){
	
	long long int x, minx = -1, ans = 0;
	while (scanf("%d", &x) != EOF) {
		if (minx == -1) minx = x;
		if ( x > minx){
			ans += x-minx;
		}
		minx = x;
	}
	cout << ans << endl;
	
	return 0;
} 
