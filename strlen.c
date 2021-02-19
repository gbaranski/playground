#include <stdio.h>
#include <stdint.h>

int main() {
  const uint16_t str[] = {
    0x48, 
    0x45,
    0x4C,
    0x4C,
    0x4F
  };
  const uint8_t arr[] = { 0xFA, 0xFF, 0xFF, 0xFF };
  uint32_t len = ~(arr[0] | (arr[1] << 8) | (arr[2] << 16) | (arr[3] << 24));

  for ( uint32_t i = 0; i < len; i++ ) {
    printf("str[%u] = 0x%08X\n", i, str[i]);
  }
}
