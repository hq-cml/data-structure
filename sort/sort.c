#include <stdio.h>
#include <stdlib.h>

void bubbleSort(int a[], int len) {
    int tmp, i, j;
    for(i=1; i < len; i++) {
        for(j=0; j < (len-i); j++) {
            if (a[j] > a[j+1]) {
                tmp = a[j];
                a[j] = a[j+1];
                a[j+1] = tmp;
            }
        }
    }
}

void foreachArr(int a[], int len) {
    int i;
    for(i=0; i<len; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

int main() {
    int a[] = {2, 3, 1, 4, 1, 6, 9, 7};
    bubbleSort(a, 8);
    foreachArr(a, 8);
}