#include <string.h>
#include <stdio.h>

void modify( char* str ) 
{
  printf("%c", str[10000]);
  str[10] = 'a';
  str = "dsahashhdsh";

}

int main(  ) {
  char* str;
  modify( str );
}
