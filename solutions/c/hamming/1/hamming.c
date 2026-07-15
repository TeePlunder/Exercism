#include "hamming.h"
#include <stdio.h>

int compute(const char *lhs, const char *rhs) {
  int count = 0;

  if (!lhs || !rhs)
    return -1;
  for (; *lhs != '\0' && *rhs != '\0'; ++lhs, ++rhs) {
    if (*lhs != *rhs)
      ++count;
  }
  if (*lhs != *rhs)
    return -1;
  return count;
}
