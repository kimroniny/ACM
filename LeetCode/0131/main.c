#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<stdbool.h>

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be calloced, assume caller calls free().
 */


void dfs(int x, int l, int pLen, int *returnSize, int *returnColumnSizes, char ***ans, char **p, char *s0, char *sr) {
    if (x >= l) {
        ++(*returnSize);
        char ** pp = (char **)malloc(sizeof(char*)*pLen);
        int i;
        for (i = 0; i < pLen; i++)  {
            int ls = strlen(*(p+i));
            pp[i] = (char *)malloc(sizeof(char)*(ls+1));
            strcpy(pp[i], *(p+i));
        }
        ans[(*returnSize)-1] = pp;
        returnColumnSizes[(*returnSize)-1] = pLen;
        return;
    }
    int i;
    for (i = x; i < l; i++) {
        if (!strncmp(s0+x, sr+l-1-i, i-x+1)) {            
            p[pLen] = (char *)malloc(sizeof(char)*(i-x+2));
            p[pLen][i-x+1] = '\0';
            strncpy(p[pLen], s0+x, i-x+1);
            dfs(i+1, l, pLen+1, returnSize, returnColumnSizes, ans, p, s0, sr);
        }
    }
}

char *** partition(char * s, int* returnSize, int** returnColumnSizes){
    int l = strlen(s);
    int MAXLEN = l * (1 << l);
    char *** ans = (char ***)malloc(sizeof(char **)*MAXLEN);
    char *s0 = NULL, *sr = NULL;
    char **p = (char **)malloc(sizeof(char*)*l);
    s0 = s;
    sr = (char *)malloc(sizeof(char)*(l+1)); sr[l] = '\0';
    int i;
    for (i = 0; i < l; i++) sr[i] = s[l-1-i];
    int pLen = 0;
    *returnSize = 0;
    *returnColumnSizes = (int *)malloc(sizeof(int)*MAXLEN);
    
    dfs(0, l, pLen, returnSize, *returnColumnSizes, ans, p, s0, sr);
    for (int i = 0; i < *returnSize; i++) {
        for (int j = 0; j < (*returnColumnSizes)[i]; j++) {
            printf("%s ", ans[i][j]);
        }
        printf("\n");
    }
    return ans;
}

void main(void) {
    
    char *s = "aab";
    int * returnSize = (int *)malloc(sizeof(int));
    int ** returnColumnSizes = (int **)malloc(sizeof(int*));
    partition(s, returnSize, returnColumnSizes);    
}