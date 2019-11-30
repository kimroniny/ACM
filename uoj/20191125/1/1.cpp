#include<iostream>
#include<cstdio>
#include<cstring>
#include<vector>

using namespace std;

typedef struct LocList{
	long long loc;
	LocList * next;
}LocList;

LocList* headm, *pm, *headb, *pb;

int main(){
	char c;
	long long cnt = 0, ans = 0;
	headm = new LocList; headm->loc = -1; headm->next = NULL; pm = headm;
	headb = new LocList; headb->loc = -1; headb->next = NULL; pb = headb;
	char* s;
	s = new char;
	gets(s);
	long long k;
	cin >> k;
	while (*s != '\0'){		
		char c = *s; 
		if (c == 'M'){
			long long loc = headb->loc;
			if (loc == -1){
				LocList * tmp = new LocList;
				tmp->loc = -1; tmp->next = NULL;
				pm->loc = cnt;
				pm->next = tmp;
				pm = tmp;
			}else{
				while (headb->next != NULL){
					if (cnt - headb->loc <= k){
						ans++;
						headb = headb->next;
						break;
					}
					headb = headb->next;
				}
			}
		}else if (c == 'B'){
			long long loc = headm->loc;
			if (loc == -1){
				LocList * tmp = new LocList;
				tmp->loc = -1; tmp->next = NULL;
				pb->loc = cnt;
				pb->next = tmp;
				pb = tmp;
			}else{
				while (headm->next != NULL){
					if (cnt - headm->loc <= k){
						ans ++;
						headm = headm->next;
						break;
					}
					headm = headm->next;
				}
			}
		}
		cnt++;
		s += 1;
	}
	cout << ans << endl;
	return 0;
}
