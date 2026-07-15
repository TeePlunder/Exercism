#include "isogram.h"
#include <string.h>

bool is_isogram(const char phrase[]) {
  int lengthOfPhrase = (int)strlen(phrase);
  int letterFoundCount = 0;
  for (int i = 0; i < lengthOfPhrase; i++) {
    char currentLetterToCheck = phrase[i];
    for (int j = 0; j < lengthOfPhrase; j++) {
      char currentLetter = phrase[j];
      if (currentLetter == currentLetterToCheck) {
        if (letterFoundCount >= 1) {
          return false;
        }
        letterFoundCount++;
      }
    }
  }
  return true;
}
