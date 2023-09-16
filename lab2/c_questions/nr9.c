#include <stdlib.h>


void function() {
    char *c = (char *) malloc(sizeof(char));
    *c = 'a';
}


int main() {
    function();
}
