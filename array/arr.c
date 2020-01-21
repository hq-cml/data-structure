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

//二维数组，从左上到右下越来越大，找寻指定数字
int find2DimArr(int **a, int rows, int cols, int n) {
    int r = 0;
    int c = cols-1;
    //将a转化为1纬数组
    //所有就有a[x][y] == *((int*)a + r*cols + c)
    while(r < rows && c >= 0) {
        while(*((int*)a + r*cols + c) > n && c >= 0)
            c--;
        if (c < 0) {
            printf("Not Find\n");
            return 0;
        }
        if (*((int*)a + r*cols + c) == n) {
            printf("Find %d in [%d, %d]\n", n, r, c);
            return 1;
        }

        if (*((int*)a + r*cols + c) < n && r < rows)
            r++;
        if (r >= rows) {
            printf("Not Find\n");
            return 0;
        }
        if (*((int*)a + r*cols + c) == n) {
            printf("Find %d in [%d, %d]\n", n, r, c);
            return 1;
        }
    }

    return 0; //没找到
}


int main() {
//    int a[] = {1, 2, 3, 4, 2, 3, 7};
//    findDuplicate(a, 7);

    int a[4][4] = {
        {1, 2, 8, 9},
        {2, 4, 9, 12},
        {5, 7, 10, 13},
        {6, 8, 11, 12},
    };
    find2DimArr((int**)a, 4, 4, 1);
    find2DimArr((int**)a, 4, 4, 6);
    find2DimArr((int**)a, 4, 4, 7);
    find2DimArr((int**)a, 4, 4, 16);
}
