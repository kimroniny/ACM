#include<iostream>
#include<cstdio>
#include<vector>
#include<queue> 
#include<cmath>
#include<utility>

using namespace std;

short m, n;
bool map[1000][1000];
queue< pair<short, short> > que;


bool check(short x, short y){
	return (0 <= x && x < m && 0 <= y && y < n);
} 

void bfs(short x, short y){
	que.push({x, y});
	while (!que.empty()){
		pair<short, short> temp = que.front();
		que.pop();
		for (short i = -1; i < 2; i++){
			for (short j = -1; j < 2; j++)
			if (abs(i) != abs(j) && 
			check(temp.first+i, temp.second+j) && 
			map[temp.first+i][temp.second+j] == 1
			){
				map[temp.first+i][temp.second+j] = 0;  // 哎呦我滴个老天啊，这里不加这句话，会重复入队啊 
				que.push({temp.first+i, temp.second+j});
			} 
		}
	}
}

int main(){
	cin >> m >> n; 
	for (short i = 0; i < m; i++){
		for (short j = 0; j < n; j++){
			cin >> map[i][j];
		}
	} 
	int ans = 0;
	for (short i = 0; i < m; i++){
		for (short j = 0; j < n; j++){
			if (map[i][j] == 1){
				ans += 1;
				bfs(i, j);
			}
		}
	}
	cout << ans << endl;
	return 0;
}
