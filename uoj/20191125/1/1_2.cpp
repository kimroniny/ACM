#include<iostream>
#include<cstdio>
#include<cmath>

using namespace std;
int m[10050], b[10050];
int cm, cb;

int main(){
	char c;
	long long cnt;
	cm = cb = 0;
	while (scanf("%c", &c) && c != '\n'){
		if (c == 'M') {
			m[cm++] = cnt++;
		}else if (c == 'B'){
			b[cb++] = cnt ++;
		} 
	}
	int k;
	cin >> k;
	int i = 0, j = 0, ans;
	while (i < cm && j < cb){
		if (abs(m[i]-b[j]) <= k){
			ans++;
			i++; j++;
		}else{
			if (m[i] < b[j]) i++;
			else j++;
		}
	}
	cout << ans << endl;
	return 0;
} 
