#include <stdio.h>
#include <stdint.h>


void fn( char *ptr ) {
  printf("0: %d\n", ptr[0]);
  char* ptr2 = ptr + 2;
  printf("0: %d\n", ptr2[0]);
}

int main() {
  char arr[20];
  for ( int i = 0; i < sizeof(arr); i++ ) {
    arr[i] = i;
  }
  printf("0: %d\n", arr[0]);
  fn(arr);

}
