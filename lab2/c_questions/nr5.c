#include <stdio.h>


int main() {
    int i = 1;
    int *p = &i;
    p += 2;
    printf("%d\n", i);
}
