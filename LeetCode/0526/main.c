#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<stdbool.h>

int p;
bool v[16];
int dfs(int x) {
    if (x == 0) return 1;
    int i;
    int k = 0;
    for (i = 1; i <= p; i++)
        if (!v[i] && (x % i == 0 || i%x == 0)) {
            v[i] = true;
            k += dfs(x-1);
            v[i] = false;
        }
    return k;
}

int countArrangement(int n )  {
    p = n;
    memset(v, 0, sizeof(v));
	return dfs(n);
}

int main(int argc, char const *argv[])
{
    printf("%d\n", countArrangement(15));
    return 0;
}
