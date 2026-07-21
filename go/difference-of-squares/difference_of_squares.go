package differenceofsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		powResult := int(math.Pow(float64(i), 2))
		sum += powResult
	}

	return sum
}

func Difference(n int) int {
	squareOfSum := SquareOfSum(n)

	sumOfSquares := SumOfSquares(n)

	return squareOfSum - sumOfSquares
}
