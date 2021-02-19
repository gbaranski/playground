#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <time.h>

int main() {
  uint8_t *arr1 = (uint8_t*)malloc(512*sizeof(uint16_t));
  for (int i = 0; i < 1024; i++) {
    arr1[i] = i % 255;
  }

  uint16_t *arr2 = (uint16_t*)malloc(512*sizeof(uint16_t));

  uint64_t *times = (uint64_t*)malloc(512*sizeof(uint64_t));

  // Fill up the arr2 with values from arr1
  for (int i = 0; i < 512; i++) {
    uint64_t start = clock();

    for (int j = 0; j < 1e5; j++) {
      arr2[i] = *(uint16_t*)&arr1[i/2];
    }

    times[i] = clock() - start;
  }
  for (int i = 0; i < 512; i++) {
    printf("%zu ", times[i]);
  }
  return 0;
}

