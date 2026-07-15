#include "armstrong_numbers.h"
#include <math.h>
#include <stdio.h>

bool is_armstrong_number(int candidate)
{
    int res = 0;
    int digit_count = log10(candidate) + 1;
    for (int i = candidate; i > 0; i /= 10)
        res += pow(i % 10, digit_count);
    return candidate == res;
}
