#include<iostream>
#include<cstdio>
#include<algorithm>
#include<vector>

using namespace std;

int n;
long long ans = 0;
vector<long long> nums;

void mergesort(int l, int r){
	if (l >= r) return;
	int mid = (l+r) >> 1;
	mergesort(l, mid);
	mergesort(mid+1, r);
	
	int i = l, j = mid+1;
//	vector<long long> temp(r+1-l);
	vector<long long> temp;
	while (i <= mid && j <= r){
		if (nums[i] <= nums[j]) {
			temp.push_back(nums[i++]);
		}else{
			if (nums[i] > 3*nums[j]){
				ans += mid+1-i;
			}else{
				if (i < mid){
					int ls = i+1, rs = mid;
					while (ls <= rs){
						int mids = (ls+rs) >> 1;
						if  (nums[mids] <= 3*nums[j]) ls = mids+1;
						else rs = mids-1;
					}
					ans += mid+1-ls;
				}
			}
			temp.push_back(nums[j++]);
		}
	}
//	if (i <= mid) copy(nums.begin()+i, nums.begin()+mid+1, temp.begin()+i-l);
//	if (j <= r) copy(nums.begin()+j, nums.begin()+r+1, temp.begin()+);
	while (i <= mid) temp.push_back(nums[i++]);
	while (j <= r) temp.push_back(nums[j++]);
	copy(temp.begin(), temp.end(), nums.begin()+l);
}

int main(){
	cin >> n;
	long long int x;
	for (int i = 0; i < n; i++)  {
		scanf("%lld", &x);
		nums.push_back(x);
	}
//	for (int i = 0; i < nums.size(); i++) cout << nums[i] << endl;
	mergesort(0, n-1);
//	for (int i = 0; i < nums.size(); i++) cout << nums[i] << endl;
	cout << ans << endl;
	return 0;
} 
