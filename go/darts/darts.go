package darts

import "math"

type TargetCircle struct {
	redius float64
	score  int
}

func isInSideTarget(targetcircle TargetCircle, x, y float64) bool {
	// calcuation for this is `(x-h)^2 + (y-k)^2 <= r `

	// h in the calculation
	circlePositionX := 0.0
	// k in the calculation
	circlePositionY := 0.0

	return math.Pow(x-circlePositionX, 2)+math.Pow(y-circlePositionY, 2) <= math.Pow(targetcircle.redius, 2)
}

func Score(x, y float64) int {
	outerCircle := TargetCircle{redius: 10, score: 1}
	middleCircle := TargetCircle{redius: 5, score: 5}
	innerCircle := TargetCircle{redius: 1, score: 10}

	if isInSideTarget(innerCircle, x, y) {
		return innerCircle.score
	}

	if isInSideTarget(middleCircle, x, y) {
		return middleCircle.score
	}

	if isInSideTarget(outerCircle, x, y) {
		return outerCircle.score
	}

	return 0
}
