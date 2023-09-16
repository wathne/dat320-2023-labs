#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[]) {
    int i = 0x7 & atoi(argv[1]);
    printf("%d \n", i);
}
