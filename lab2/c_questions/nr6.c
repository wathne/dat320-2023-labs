#include <stdio.h>


void set_val(int i, int new_val) {
    i = new_val;
}


int main() {
    int i = 2;
    set_val(i, 3);
    printf("%d\n", i);
}
