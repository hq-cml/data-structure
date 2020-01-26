#include <stdio.h>
#include <stdlib.h>

#define EPS 0.000001

int EqualZero(float f) {
    return (-EPS <= f && f <= EPS);
}

float power(float base, int n) {
    int flag = 0;
    float ret = 1.0;
    if(EqualZero(base)) {
        return 0.0;
    }

    if(n<0) {
        flag = 1;
        n = n*(-1);
    }

    while(n--) {
        ret = ret * base;
    }

    if (flag) {
        return 1/ret;
    }
    return ret;
}

int main() {
    printf("%f\n", power(5, 2));
    printf("%f\n", power(5, -2));
}