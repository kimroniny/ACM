#include<algorithm>
#include<iostream>
#include<cstdio>
#include<cstdlib>
#include<vector>

using namespace std;

bool cmp(const int& a, const int& b){
	return a < b;
}

int main() {
    int n, k, x;
    vector<int> nums;
    cin >> n >> k;
    for (int i = 0; i < n; i++){
    	cin >> x;
    	nums.push_back(x);
	}
	sort(nums.begin(), nums.end(), cmp);
	cout << nums[n-k] << endl;
    return 0;
}
