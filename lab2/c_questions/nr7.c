#include <stdlib.h>


typedef struct process
{
    int pid;
} process_t;


int main() {
    process_t *p = (process_t *) malloc(sizeof(process_t));
    (*p).pid = 10;
}
