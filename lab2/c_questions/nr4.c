#include <stdio.h>


void set_val(int *i, int new_val) {
    *i = new_val;
}


int main() {
    int i = 10;
    int *p = &i;
    set_val(p, 30);
    printf("%d\n", i);
}
