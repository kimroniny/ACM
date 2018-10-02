#include<iostream>
#include<cstdio>
using namespace std;
// double a[]
int main(){
	double x;
	while (scanf("%lf", &x) != EOF && x != 0.00){
		int i = 2;
		double sum = 0;
		while (sum < x){
			sum += 1.0/i;
			i++;
		}
		cout << i-2 << " card(s)" << endl;
		
	}
	return 0;
} 
