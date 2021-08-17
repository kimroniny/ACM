#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<stdbool.h>

void test1() {
    char **p = malloc(sizeof(char*)*2);
    char *s = malloc(sizeof(char)*3);
    s[0] = '1';
    s[1] = '2';
    s[2] = '\0';
    p[0] = s;
    printf("%d\n%p\n%s\n%s\n%c\n%d\n%d", strlen(*p), *p, *p, s, *(*p + 1), strlen(s), sizeof(char **));
}

void test2() {
    char *s = "0123456";
    printf("s.length = %d\n", strlen(s));
}

void test3() {
    int* p = malloc(sizeof(int) * 5);
    printf("p.length = %d\n", sizeof(p)/sizeof(int));
}

int* test4func(int *p) {
    p = (int *)realloc(p, sizeof(int) * 100000);
    p[0] = 3;
    return p;
}

void test4() {
    int *p = (int *)malloc(sizeof(int));
    *p = 2;
    int *pp = test4func(p);
    printf("%d, %d\n", p[0], pp[0]);
}

int main(int argc, char const *argv[])
{
    // test1();
    test4();
    return 0;
}
