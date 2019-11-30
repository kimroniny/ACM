#include<iostream>
#include<cstdio>
#include<algorithm>
#include<vector>

using namespace std;

int n;
vector<long long> nums;

int main(){
	cin >> n;
	long long int x;
	for (int i = 0; i < n; i++)  {
		scanf("%lld", &x);
		nums.push_back(x);
	}
//	for (int i = 0; i < nums.size(); i++) cout << nums[i] << endl;
//	mergesort(0, n-1);
	vector<long long> temp(nums.size());
	copy(nums.begin(), nums.end(), temp.begin());
	cout << temp.size() << endl;
	for (int i = 0; i < temp.size(); i++) cout << temp[i] << endl;
	return 0;
} 
