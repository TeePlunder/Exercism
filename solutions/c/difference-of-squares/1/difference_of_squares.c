#include "difference_of_squares.h"
#include <math.h>

unsigned int sum_of_squares(unsigned int number) {
  unsigned int sum = 0;

  for (unsigned int i = 0; i <= number; i++) {
    sum += pow(i, 2);
  }

  return sum;
}

unsigned int square_of_sum(unsigned int number) {
  unsigned int sum = 0;

  for (unsigned int i = 0; i <= number; i++) {
    sum += i;
  }

  return pow(sum, 2);
}

unsigned int difference_of_squares(unsigned int number) {
  return square_of_sum(number) - sum_of_squares(number);
}
