#include <stdio.h>
#include <stdint.h>
#include <string.h>

int main(void) {
  uint8_t a[] = {0, 0, 0, 0x3};

  uint32_t b = a[3] | (a[2] << 8) | (a[1] << 16) | (a[0] << 24);
  printf("b: %d", b);

  
  return 0;
}
