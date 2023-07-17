#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int x, int y)
{
    return x > y ? x : y;
}

char *addStrings(char *num1, char *num2)
{
    int l1 = strlen(num1) - 1;
    int l2 = strlen(num2) - 1;
    int lp = max(l1, l2) + 1;
    int i = l1, j = l2, k = lp;
    int x = 0, y = 0, jin = 0;
    char *p = (char *)malloc(sizeof(char) * (lp + 2));
    p[k + 1] = '\0';

    while (i >= 0 || j >= 0)
    {
        if (j < 0)
        {
            y = 0;
        }
        else
        {
            y = *(num2 + j) - '0';
            j -= 1;
        }
        if (i < 0)
        {
            x = 0;
        }
        else
        {
            x = (*(num1 + i)) - '0';
            i -= 1;
        }
        p[k] = '0' + (x + y + jin) % 10;
        jin = (x + y + jin) / 10;
        k -= 1;
    }
    if (jin != 0)
    {
        p[k] = '0' + jin;
    }
    else
    {
        k += 1;
    }
    p = p + k;
    return p;
}

int main(void)
{
    char *s1 = "11";
    char *s2 = "9";
    char *p = addStrings(s1, s2);
    printf("%s\n", p);
    return 0;
}