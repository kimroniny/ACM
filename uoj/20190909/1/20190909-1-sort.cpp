#include<iostream>
#include<algorithm>
#include<vector>
#include<cstdio>

using namespace std;
 
int main(){
	int n, k, x;
	cin >> n >> k;
	vector<int> nums;
	for (int i = 0; i < n; i++){
		scanf("%d", &x);  // use scanf to read lots of data
		nums.push_back(x);
	}
	
	int low = 0, high = n-1;
	while(low <= high){
		int i = low, j = high;
		while(i <= j){
			while (i <= j && nums[i] >= nums[low]) i++;
			while (i <= j && nums[j] < nums[low]) j--;
			if (i < j) swap(nums[i++], nums[j--]);
		}
		swap(nums[j], nums[low]);
		if (j == k-1){
			cout << nums[j] << endl;
			break;
		}
		if (j < k-1) low = j+1; else high = j-1;
	} 
	
	return 0;
} 
