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

//快速排序
void quickSort(int a[], int low, int high) {
    if (low >= high) {
        return;
    }
    int mid = partition(a, low, high);
    quickSort(a, mid+1, high);   //recursion
    quickSort(a, low, mid-1);
}
//5 6 2 1 8 10 9  - 5
int partition(int a[], int low, int high) {
    int val = a[low];
    int tmp;
    while(low < high) {
        while (low < high && val <= a[high]) {
            high --;
        }
        a[low] = a[high];
        while (low < high && a[low] <= val) {
            low ++;
        }
        a[high] = a[low];
    }
    a[low] = val;
    return low;
}

void foreachArr(int a[], int len) {
    int i;
    for(i=0; i<len; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

int main() {
    //int a[] = {2, 3, 1, 4, 1, 6, 9, 7};
    //bubbleSort(a, 8);

    int a[] = {5, 6, 2, 1, 8, 10, 9};
    quickSort(a, 0, 6);
    foreachArr(a, 7);
}