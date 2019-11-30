#include<iostream>
#include<algorithm>
#include<vector>

using namespace std;

int main(){
	int n, x;
	vector<int> nums;
	cin >> n;
	for (int i = 0; i < n; i ++){
		cin >> x;
		nums.push_back(x);
	}
	auto cmp = [](const int& a, const int& b) {
		return a > b;
	};
	make_heap(nums.begin(), nums.end(), cmp);
	
	for (int i = 0; i < n; i++){
		cout << "front: " << nums.front() << endl;
		pop_heap(nums.begin(), nums.end(), cmp);
		cout << "back: " << nums.back() << endl;
		nums.pop_back();
	}
	return 0;
}
