#include "isogram.h"
#include <string.h>
#include <ctype.h>

bool is_isogram(const char phrase[]) {
  if (phrase == NULL) {
    return false;
  }
  int lengthOfPhrase = (int)strlen(phrase);
  int letterFoundCount = 0;
  for (int i = 0; i < lengthOfPhrase; i++) {
    char currentLetterToCheck = tolower(phrase[i]);
    int currentLetterIndex = i;
    for (int j = 0; j < lengthOfPhrase; j++) {
      char currentLetter = tolower(phrase[j]);
      if(currentLetterIndex == j || !isalpha(currentLetter)){
        continue;
      }
      if (currentLetter == currentLetterToCheck) {
        if (letterFoundCount > 0) {
          return false;
        }
        letterFoundCount++;
      }
    }
  }
  return true;
}
