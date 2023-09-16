#include <stdio.h>
#include <stdlib.h>


int main(int argc, char* argv[]) {
    int i = 2;
    if (argc >= 4) {
        i = i * atoi(argv[1]) * atoi(argv[2]) * atoi(argv[3]);
    }

    printf("%d \n", i);
}
