#include <stdio.h>
#include <string.h>

int main() {
    if (strcmp("abc", "abc")) {
        printf("equal");
    } else {
        printf("not equal");
    }
    return 0;
}
