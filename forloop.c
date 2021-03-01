#include <stdio.h>

int main( void ) {
  int n;
  scanf( "%d", &n );
  for ( int i = 0; i < n; i++ ) {
    printf( "%d:Dzien dobry\n", i + 1 );
  }

  return 0;
}
