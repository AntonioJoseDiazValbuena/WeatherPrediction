package utils

import "math"

func FindSlope(x1, x2, y1, y2 float64) float64 {
	dX := x2 - x1
	dY := y2 - y1

	return Round(dY / dX)
}

func Round(value float64) float64 {
	return math.Round(value*100) / 100
}
