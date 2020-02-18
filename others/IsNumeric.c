#include <stdio.h>
#include <stdlib.h>

int scanUnsignedInt(const char **p) {
    char *before = *p;
    while(**p != '\0' && **p >= '0'  && **p <= '9' )
        (*p)++;
    return *p > before;
}

int scanSignedInt(const char **p) {
    if(**p == '+' || **p == '-') {
        (*p)++;
    }
    return scanUnsignedInt(p);
}

int IsNumeric(const char *p) {
    if (p == NULL) {
        return 0;
    }

    int ret = scanSignedInt(&p);
    if (*p == '.') {
        //小数点
        p ++;
        ret = scanUnsignedInt(&p) || ret; //用|| 因为整数部分可以不存在, 并且scanUnsignedInt需要放在前面先做
    }

    if (*p == 'e' || *p == 'E') {
        //存在指数部分
        p ++;
        ret = ret && scanSignedInt(&p);
    }
    return ret && *p == '\0'; //必须达到了字符串末尾
}

int main() {
    printf("%d\n", IsNumeric("+100"));
    printf("%d\n", IsNumeric("5e2"));
    printf("%d\n", IsNumeric("3.14"));
    printf("%d\n", IsNumeric("-1E-10"));

    printf("%d\n", IsNumeric("12e"));
    printf("%d\n", IsNumeric("1a3.14"));
    printf("%d\n", IsNumeric("1.1.1"));
    printf("%d\n", IsNumeric("11e+2.1"));
}