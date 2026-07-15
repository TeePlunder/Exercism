#include "square_root.h"
#include <stdint.h>

uint16_t square_root(uint16_t n)
{
  uint16_t lower = 1;
  uint16_t upper = 256;
  while(1) {
    uint16_t mid = (upper + lower) / 2;
    uint16_t square = mid * mid;
    if(square < n)
      lower = mid;
    else if(square > n)
      upper = mid;
    else
      return mid;
  }

}
