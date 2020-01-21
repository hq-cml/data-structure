#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//a[n]中，数字只会是1-n, 重复的数字
int findDuplicate(unsigned int a[], unsigned int len) {
    int *p = (int*)malloc(sizeof(int) * (len+1));
    if(!p) {
        return -1;
    }
    memset(p, 0, sizeof(int) * (len+1));
    int i;
    for(i=0; i<len; i++) {
        p[a[i]] ++;
    }

    for (i=0; i<=len; i++) {
        if(p[i] > 1) {
            printf("%d\n", i);
        }
    }
}

int main() {
    int a[] = {1, 2, 3, 4, 2, 3, 7};
    findDuplicate(a, 7);
}
