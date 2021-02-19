#include <stdio.h>
#include <time.h>


int main() {
  int timestamp = time(NULL);
  printf("timestamp: %d\n", timestamp);
}
