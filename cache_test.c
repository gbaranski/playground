#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <x86intrin.h>

int main() {
  volatile uint8_t *arr1 = malloc(4096 * 16);
  for (size_t i = 0; i < 4096 * 16; i++) {
    arr1[i] = rand();
  }

  printf("%p\n", arr1);

  static uint32_t res[4096];
  static uint32_t resx[4096];

  for (int i = 0; i < 256; i++) {

    for (int k = 0; k < 20; k++) {  // Repeat test 20 times, select best time.

      uint16_t result = 0;
      uint64_t start = _rdtsc();

      for (int j = 0; j < 1e4; j++) {
        _mm_mfence();  // Memory fence.
        result ^= *(volatile uint16_t*)(arr1+i) | result;
        _mm_mfence();  // Memory fence.
      }

      uint32_t t = (uint32_t)(_rdtsc() - start);

      if (res[i] == 0 || t < res[i]) {
        res[i] = t;
      }

      resx[i] = result;
    }

  }

  for (int i = 0; i < 256; i++) {
     printf("...%.3x: %u\n", ((uint32_t)(uint64_t)(&arr1[i])) & 0xfff, res[i], resx[i]);
  }


  return 0;
}
