#include <cstdio>
#include <cmath>
int main()
{
    int n;
    scanf("%d",&n);
    while(n--){
        double a,b,c;
        double pd;
        double part1,part2;
        scanf("%lf %lf %lf",&a,&b,&c);
        if (a < 0){
                    a = 0 - a;
                    b = 0 - b;
                    c = 0 - c;
                }
        pd=b*b-4*a*c;
        part1 = (0-b)/ (2 * a);
        if(pd>0){
            part2 = sqrt(pd)/(2 * a);
            printf("x1=%.5lf;x2=%.5lf\n",part1+part2,part1-part2);
        }else if(fabs(pd)<1e-8){
            printf("x1=x2=%.5lf\n",part1);
        }else{
            part2 = sqrt(-pd)/ (2 * a);
            printf("x1=%.5lf+%.5lfi;x2=%.5lf-%.5lfi\n",part1,part2,part1,part2);
        }
    }
    return 0;
}
