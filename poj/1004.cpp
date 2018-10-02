#include<iostream>
using namespace std;
int main(){
	double x, ans = 0;
	while(cin >> x){
		ans += x;
	}
	cout << "$" << ans/12.0 << endl;
	return 0;
}
