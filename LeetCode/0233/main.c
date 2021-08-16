#include<stdio.h>
#include<stdlib.h>
#include<math.h>

int countDigitOne(int n){
    int cnt = log10(n), n2 = n;
    long long sum = 0;
    int i; 
    long long ans = 0, p = 1, d=0;
    for (i = 0; i <= cnt; i++) {
        int x = n2 % 10, k = 0;
        k = x>1?p:(sum+1)*x;
        ans += x*d + k;
        d = 10*d+p;
        sum += x*p;
        p *= 10;
        n2 /= 10;
    }
    return ans;
}

int main(void) {
    printf("%d\n", countDigitOne(101));
    return 0;
}