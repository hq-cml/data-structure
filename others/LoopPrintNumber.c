#include <stdio.h>
#include <stdlib.h>
# include <string.h>

void PrintNumber(char *p, int len) {
    int i = 0;
    for(i=0; i<len; i++) {
        if (p[i] != '0') {
            printf("%s\n", p+i);
            break;
        }
    }
}

void LoopPrintNumber(int n) {
    char *p = (char *)malloc(sizeof(char) * (n+1));
    if (!p) {
        return;
    }

    //初始化
    int end = 0;
    int i=0;
    //memset(p, '0', sizeof(char) * (n-1));
    for (i=0; i<n; i++) {
        p[i] = '0';
    }
    p[n] = 0;
    while(!end) {
        p[n-1] ++; //自增
        if (p[n-1] > '9') {
            //存在进位问题
            for(i=n-1; i>=0; i--) {
                if (p[i] > '9') {
                    if (i==0) {
                        end = 1;
                        break;
                    }
                    p[i] = '0';
                    p[i-1]++;
                }
            }
        }
        if (!end) {
            PrintNumber(p, n);
        }
    }
}

int main() {
    LoopPrintNumber(3);
}