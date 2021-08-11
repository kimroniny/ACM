#include<stdio.h>
#include<stdlib.h>
#include<uthash.h>
#include<src/uthash.h>

typedef struct my_struct {
    int id;                    /* key */
    char name[10];
    UT_hash_handle hh;         /* makes this structure hashable */
}MyStruct;

int numberOfArithmeticSlices(int* nums, int numsSize){
    MyStruct *infos = NULL;
    
    
}

int main(int argc, char const *argv[])
{
    int num[1050];
    
    printf("%d\n", numberOfArithmeticSlices(
        num,
        sizeof(num)/sizeof(num[0])
    ));
    return 0;
}
