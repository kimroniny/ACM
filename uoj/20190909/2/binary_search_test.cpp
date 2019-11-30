#include<iostream>
#include<cstdio>
#include<vector>
#include<algorithm>

using namespace std;

int main(){
	int n, x, k;
	cin >> n;
	vector<int> nums;
	for (int i = 0; i < n; i++) {
		cin >> x; nums.push_back(x);
	}
	sort(nums.begin(), nums.end());
	cin >> k;
	int l = 0, r = n-1;
	while (l <= r){
		int mid = (l+r)>>1;
		if (nums[mid] <= k) l = mid+1;
		else r = mid-1;
	}
	cout << r << endl;
	cout << l << endl;
	return 0;
}
