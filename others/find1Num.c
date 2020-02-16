#include <stdio.h>
#include <stdlib.h>

int find1NumV1(int n) {
    int flag = 1;
    int cnt = 0;
    while(flag) {
        if (flag & n) cnt ++;
        flag = flag<<1;
    }
    return cnt;
}

int find1NumV2(int n) {
    int flag = 1;
    int cnt = 0;
    while(n) {
        cnt ++;
        n = n & (n-1);
    }
    return cnt;
}

int main() {
    printf("%d\n", find1NumV1(7));
    printf("%d\n", find1NumV1(-1));
    printf("%d\n", find1NumV2(7));
    printf("%d\n", find1NumV2(-1));
}