#include<iostream>
#include<cstdio>

using namespace std;

int main(){
	int n;
	long long x;
	cin >> n;
	int count = 0, p;
	for (int i = 0; i < n; i++){
		scanf("%lld", &x);
		if (count == 0) {
			count = 1;
			p = x;
		}else{
			if (x == p) count++;
			else count--;
		}
	} 
	cout << p << endl; 
	return 0;
}
 
