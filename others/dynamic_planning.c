#include <stdio.h>
#include <stdlib.h>

//贪心算法
int maxProductAfterCutting(int length)
{
    if(length < 2)
        return 0;
    if(length == 2)
        return 1;
    if(length == 3)
        return 2;

    int* products = (int*)malloc(sizeof(int) * (length+1));
    //当length小于4的时候，是比较特殊的情况，f(n)的最优值就是n，因为减了不如不减
    //因为必须要剪一刀，所以，上面有一些特殊情况的枚举，但是当n>4的时候，这些子情况最优值是不减
    products[0] = 0;
    products[1] = 1;
    products[2] = 2;
    products[3] = 3;

    int max = 0;
    int i,j;
    for(i = 4; i <= length; ++i)
    {
        max = 0;
        for(j = 1; j <i; ++j)
        {
            int product = products[j] * products[i - j];
            if(max < product)
                max = product;

            products[i] = max;
        }
    }

    max = products[length];
    free(products);

    return max;
}

int main() {
    printf("%d\n", maxProductAfterCutting(2));
    printf("%d\n", maxProductAfterCutting(3));
    printf("%d\n", maxProductAfterCutting(4));
    printf("%d\n", maxProductAfterCutting(8));
}