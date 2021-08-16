#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<stdbool.h>

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
char *** ans = NULL;
char *s0 = NULL, *sr = NULL;
char **p = NULL;
int l;

void dfs(int x) {
    if (x >= l) {
        int k = sizeof(ans)/sizeof(char*);
        ans = (char ***)realloc(ans, sizeof(char*)+sizeof(ans));
        printf("ans.length=%d\n", sizeof(ans)/sizeof(char *));
        char ** pp = (char **)malloc(sizeof(p));
        int lp = sizeof(p)/sizeof(char *);
        printf("p.length=%d\n", lp);
        int i;
        for (i = 0; i < lp; i++)  {
            int ls = strlen(*(p+i));
            printf("ls=%d, p+i=%s\n", ls, *(p+i));
            char *s = (char *)malloc(sizeof(char) * (ls+1));
            s[ls] = '\0';
            strcpy(s, *(p+i));

            pp[i] = s;
            printf("s=%s\n", s);
            printf("pp[%d]=%s\n", i, pp[i]);
        }
        
        ans[k-1] = pp;
        return;
    }
    int i;
    for (i = x; i < l; i++) {
        if (!strncmp(s0+x, sr+l-1-i, i-x+1)) {
            int lp = sizeof(p)/sizeof(char*);
            printf("lp=%d\nsizeof(p)=%d\n", lp, sizeof(p));
            p = (char **)realloc(p, sizeof(p)+sizeof(char *));
            p[lp] = (char *)malloc(sizeof(char)*(i-x+2));
            p[lp][i-x+1] = '\0';
            strncpy(p[lp], s0+x, i-x+1);
            printf("p[lp]=%s\n", p[lp]);
            dfs(i+1);
            free(p+sizeof(p)/sizeof(char *)-1);
        }
    }
}

char *** partition(char * s, int* returnSize, int** returnColumnSizes){
    l = strlen(s);
    s0 = s;
    sr = (char *)malloc(sizeof(char)*(l+1)); sr[l] = '\0';
    int i;
    for (i = 0; i < l; i++) sr[i] = s[l-1-i];
    dfs(0);
}

void main(void) {
    char *s = "aab";
    int * returnSize = NULL;
    int ** returnColumnSizes = NULL;
    partition(s, returnSize, returnColumnSizes);
    int i, j;
    for (i = 0; i < sizeof(ans)/sizeof(char *); i++) {
        for (j = 0; j < sizeof(ans[i])/sizeof(char *); i++) {
            printf("%s ", ans[i][j]);
        }
        printf("\n");
    }
}