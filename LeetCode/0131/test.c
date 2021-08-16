#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<stdbool.h>

int main(int argc, char const *argv[])
{
    char **p = malloc(sizeof(char*)*2);
    char *s = malloc(sizeof(char)*3);
    s[0] = '1';
    s[1] = '2';
    s[2] = '\0';
    p[0] = s;
    printf("%d\n%p\n%s\n%s\n%c\n%d\n%d", strlen(*p), *p, *p, s, *(*p + 1), strlen(s), sizeof(char **));
    return 0;
}
