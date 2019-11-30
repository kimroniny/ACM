#include<iostream>
#include<cstdio>
#include<cstring>
#include<vector>

using namespace std;

int m[10050], b[10050];
int cm, cb;

typedef struct LocList{
	long long loc;
	LocList * next;
}LocList;

LocList* headm, *pm, *headb, *pb;
char alpha[2] = {'B', 'M'};

int pointer(char *s, int k){
	char c;
	long long cnt = 0, ans = 0;
	headm = new LocList; headm->loc = -1; headm->next = NULL; pm = headm;
	headb = new LocList; headb->loc = -1; headb->next = NULL; pb = headb;
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
	return ans;
}

int vect(char *s, int k){
	char c;
	long long cnt = 0;
	cm = cb = 0;
	while ((*s != '\0')){
		c = *s;
		if (c == 'M') {
			m[cm++] = cnt++;
		}else if (c == 'B'){
			b[cb++] = cnt ++;
		} 
		s+=1;
	}
	int i = 0, j = 0, ans = 0;
	while (i < cm && j < cb){
		if (abs(m[i]-b[j]) <= k){
			ans++;
			i++; j++;
		}else{
			if (m[i] < b[j]) i++;
			else j++;
		}
	}
	return ans;
}

int main(){
	int size = 4000;
	for (int i = 0; i < 500; i++){
		char *s = new char[size+1];
		for (int j = 0; j < size; j++){
			s[j] = alpha[rand()%2];
		}
		s[size] = '\0';
		
		int k = 1;
//		printf("%s\n%d\n", s, k);
		int p = pointer(s, k), v = vect(s, k);
		printf("pointer: %05d, vect: %05d ==== %d\n", p, v, p==v); 
	}
	return 0;
}
