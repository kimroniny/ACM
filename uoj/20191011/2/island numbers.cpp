#include<iostream>
#include<cstdio>
#include<vector>
#include<cmath>

using namespace std;

int m, n;
bool map[1005][1005];

bool check(short x, short y){
	if (0 <= x && x < m && 0 <= y && y < n) return true;
	return false;
} 

void dfs(short x, short y){
	map[x][y] = 0;
	for (short i = -1; i < 2; i++){
		for (short j = -1; j < 2; j++)
		if (abs(i) != abs(j) && check(x+i, y+j) && map[x+i][y+j] == 1){
			dfs(x+i, y+j);
		} 
	}
}

int main(){
	cin >> m >> n; 
	for (int i = 0; i < m; i++){
		for (int j = 0; j < n; j++){
			scanf("%d",&map[i][j]);
		}
	} 
	int ans = 0;
	for (short i = 0; i < m; i++){
		for (short j = 0; j < n; j++){
			if (map[i][j] == 1){
				ans += 1; 
				dfs(i, j);
			}
		}
	}
	cout << ans << endl;
	return 0;
}
