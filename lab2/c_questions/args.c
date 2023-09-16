#include <stdio.h>


int main(int argc, char *argv[]) {
    if (argc >= 3) {
        switch (*argv[2]) {
            case 'a':
                printf("a \n");
                break;
            case 'b':
                printf("b \n");
                break;
            case 'c':
                printf("c \n");
                break;
            case 'd':
                printf("d \n");
                break;
            default:
                printf("No match \n");
                break;
        }
    }
}
