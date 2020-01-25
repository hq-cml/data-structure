#include <stdio.h>
#include <stdlib.h>

//递归
int Fibonacci_Recurse(int n) {
    if (n<=0) {
        return 0;
    }
    if (n==1) {
        return 1;
    }
    return Fibonacci_Recurse(n-2) + Fibonacci_Recurse(n-1);
}

//非递归
int Fibonacci_NoRecurse(int n) {
    if (n<=0) {
        return 0;
    }
    if (n==1) {
        return 1;
    }

    int i=0;
    int pre = 0;
    int curr = 1;
    int total = 0;
    for(i=1; i<n; i++) {
        total = curr + pre;
        pre = curr;
        curr = total;
    }
    return total;
}

int main() {
//    printf("%d\n", Fibonacci_Recurse(0));
//    printf("%d\n", Fibonacci_Recurse(1));
//    printf("%d\n", Fibonacci_Recurse(2));
//    printf("%d\n", Fibonacci_Recurse(3));
//    printf("%d\n", Fibonacci_Recurse(20));
//    printf("%d\n", Fibonacci_Recurse(100)); //阻塞

      printf("%d\n", Fibonacci_NoRecurse(0));
      printf("%d\n", Fibonacci_NoRecurse(1));
      printf("%d\n", Fibonacci_NoRecurse(2));
      printf("%d\n", Fibonacci_NoRecurse(3));
      printf("%d\n", Fibonacci_NoRecurse(20));
      printf("%d\n", Fibonacci_NoRecurse(100));
}