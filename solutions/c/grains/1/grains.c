#include "grains.h"
#include <stdint.h>
#include <stdio.h>

uint64_t square(uint8_t index) {
  if (index < 1 || index > 64) {
    return 0ULL;
  }

  uint64_t value = 1;
  return value << (index - 1);
}

uint64_t total(void) {
  uint64_t totalAmount = 0;
  uint64_t one = 1;
  for (uint8_t i = 0; i < 64; i++) {
    totalAmount += one << i;
  }
  return totalAmount;
}
