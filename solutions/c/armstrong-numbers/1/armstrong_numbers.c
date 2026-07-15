#include "armstrong_numbers.h"
#include <math.h>

int getDigitInInt(int givenInt, int position) {
  while (--position) {
    givenInt /= 10;
  }
  return givenInt % 10;
}

bool is_armstrong_number(int toCheckNumber) {
  if (toCheckNumber < 10) {
    return true;
  }

  int digitCount = log10(toCheckNumber) + 1;
  int total = 0;

  for (int i = 0; i < digitCount; i++) {
    int currentDigit = getDigitInInt(toCheckNumber, i + 1);
    total += pow(currentDigit, digitCount);
  }

  return toCheckNumber == total;
}
