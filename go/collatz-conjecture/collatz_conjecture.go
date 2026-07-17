package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	currentNumber := n
	steps := 0

	if currentNumber <= 0 {
		return 0, errors.New("input must be a positive integer")
	}

	for currentNumber != 1 {
		if currentNumber%2 == 0 {
			currentNumber /= 2
		} else {
			currentNumber = currentNumber*3 + 1
		}
		steps++
	}

	return steps, nil
}
