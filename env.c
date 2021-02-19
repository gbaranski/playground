#include <stdio.h>

#ifndef TERM
  #define TERM "Unknown"
#endif


int main() {
  printf("term: %s\n", TERM);
  return 0;
}
