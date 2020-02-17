#include <stdio.h>
#include <stdlib.h>

int Match(char *str, char *pattern) {
    if(str == NULL || pattern == NULL) {
        return 0;
    }

    if(*str == '\0' && *pattern == '\0') {
        return 1;
    }

    if(*str != '\0' && *pattern == '\0') {
        return 0;
    }

    if(*(pattern + 1) == '*') {
        if(*pattern == *str || (*pattern == '.' && *str != '\0')) {
            //匹配上了，可以有多重选择
            return  Match(str + 1, pattern + 2)     //匹配一个
                    || Match(str + 1, pattern + 1)  //匹配两个
                    || Match(str, pattern + 2);      //没匹配忽略
        } else {
            //没匹配上，则将*作为匹配0次来处理
            return Match(str, pattern + 2);
        }
    }

    //非*，处理相对简答
    if (*str == *pattern || (*pattern == '.' && *str != '\0')) {
        return Match(str + 1, pattern + 1);
    }

    //第一个直接匹配失败了
    return 0;
}

int main() {
    printf("%d\n", Match("aaa", "a.a"));
    printf("%d\n", Match("aaa", "ab*ac*a"));
    printf("%d\n", Match("aaa", "ab*a"));
}