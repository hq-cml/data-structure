#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//模拟urlencode，将空格进行替换%20，要求在同一个字符串进行
void convertWhite(char *s, int totalLen) {
    int strLen = strlen(s);
    int whiteCnt = 0;
    int i;
    for(i=0; i<strLen; i++) {
        if(s[i] == ' ') {
            whiteCnt ++;
        }
    }

    int newLen = strLen + whiteCnt * 2;
    if ((newLen+1) > totalLen) {
        printf("No enough!\n");
    }

    char *p1 = s + strLen;
    char *p2 = s + newLen;
    while((strLen+1)>0 && p1 != p2) { //替换个数是strLen+1，因为'\0'
        if(*p1 != ' ') {
            *p2 -- = *p1--;
        } else {
            p1--;
            *p2-- = '0';
            *p2-- = '2';
            *p2-- = '%';
        }
        strLen --;
    }
}

int main() {
    char *s = malloc(sizeof(char) * 100);
    if (s == NULL) {
        return;
    }
    memset(s, 0, 100);

    //strcpy(s, "we are people!");
    //strcpy(s, "we are people! ");
    //strcpy(s, "wearepeople!");
    //strcpy(s, "  ");
    strcpy(s, " we are people! ");
    printf("ORG:%s\n", s);
    convertWhite(s, 100);
    printf("DST:%s\n", s);
}