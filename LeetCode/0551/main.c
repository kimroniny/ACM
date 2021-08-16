#include<stdio.h>
#include<string.h>
#include<stdlib.h>
#include<stdbool.h>

bool checkRecord(char * s){
    int i;
    int cntA = 0, cntL = 0;
    for (i=0;i<strlen(s);i++){
        switch (*(s+i))
        {
        case 'A':
            if (++cntA >= 2) return false;
            cntL = 0;
            break;
        case 'L':
            if (++cntL >= 3) return false;
            break;
        case 'F':
            cntL = 0;
            break;
        default:
            cntL = 0;
            break;
        }
    }
    return true;
}

void main(void) {
    char *s = "PPALLPLA";
    printf("%d\n", checkRecord(s));
}