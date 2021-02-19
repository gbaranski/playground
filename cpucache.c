#include <stdio.h>
#include <stdint.h>
#include <time.h>

int main() {
  uint8_t arr1[1024];
  uint16_t arr2[sizeof(arr1)/2];

  uint64_t times[sizeof(arr2)];

  for (int i = 0; i < 1024; i+=2) {
    uint64_t start = clock();

    for (int j = 0; j < 1e5; j++) {
      arr2[i/2] = (uint16_t)arr1[i];
    }

    times[i/2] = clock() - start;
  }
  for (int i = 0; i < 512; i++) {
    printf("%zd ", times[i]);
  }
  return 0;
}
