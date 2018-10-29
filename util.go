package raytracer

import (
        "math"
)

func FloatEqual(a, b float64) bool {
	epsilon := .00001
	return (math.Abs(a - b) < epsilon)
}

func IsPositive(f float64) bool {
	return math.Abs(f) == f
}

func Clamp(num, min, max float64) float64 {
	if num < min {
		return min
	} else if num > max {
		return max
	} else {
		return num
	}
}
