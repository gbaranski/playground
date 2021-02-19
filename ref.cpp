#include <stdio.h>

void fn(int &a) {
  a = 5;
  
}

int main() {
  int a = 5;
  fn(a);

  return 1;
}
