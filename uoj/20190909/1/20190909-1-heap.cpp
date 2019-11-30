#include<iostream>
#include<vector>
#include<algorithm>
#include<cstdio>

using namespace std;

bool cmp(const int& a, const int& b){
	return a > b;
}

int main(){
	int n, k, x;
	cin >> n >> k;
	vector<int> nums;
	for (int i = 0; i < n; i ++){
		scanf("%d", &x);
		if (i < k) {
			nums.push_back(x);
			if (i == k-1) make_heap(nums.begin(), nums.end(), cmp);
		}else{
			if (x >= nums.front()){
				pop_heap(nums.begin(), nums.end(), cmp);
				nums.pop_back();
				nums.push_back(x);
				push_heap(nums.begin(), nums.end(), cmp); 
			}
		}
	} 
	cout << nums.front() << endl; 
	
	return 0;
}
