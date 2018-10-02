#include<iostream>
#include<cstdio>
#include<cmath>
using namespace std;
const double PI = 3.1415926;
int main(){
	int n;
	cin >> n;
	for (int i = 0; i < n; i++){
		double x, y;
		scanf("%lf%lf", &x, &y);
		double ans = (x*x + y*y) * PI / 2 / 50;
		printf("Property %d: This property will begin eroding in year %d.\n", i+1, int(ceil(ans)));
	}
	cout << "END OF OUTPUT." << endl;
	return 0;
}


