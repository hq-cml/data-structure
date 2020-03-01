#include <stdio.h>
#include <stdlib.h>

int gError;
int Atoi(const char *str) {
    gError = 0;
    if (str == NULL) {
        gError = 1;
        return 0;
    }

    int positive = 1;
    int num = 0;
    const char *p = str;
    if (*p == '+') {
        p ++;
    } else if (*p == '-') {
        positive = -1;
        p ++;
    }

    if (*p == '\0') {
        gError = 1;
        return 0;
    }

    while(*p != '\0') {
        if (*p > '9' || *p < '0') {
            break;
        }
        num = num * 10 + (*p - '0');
        p ++;
    }

    if (*p != '\0') {
        gError = 1;
        return 0;
    }

    return positive * num;
}

int main() {
    //printf("%d, %d\n", Atoi("+100"), gError);
    //printf("%d, %d\n", Atoi("-108"), gError);
    //printf("%d, %d\n", Atoi("-1a8"), gError);
    printf("%d, %d, %p\n", Atoi("-"), gError, &gError);
}